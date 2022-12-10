package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileEntityType string

var (
	Root    FileEntityType = "root"
	DirType FileEntityType = "dir"
	File    FileEntityType = "file"
)

type InstructionType string

var (
	Ls  InstructionType = "ls"
	Cd  InstructionType = "cd"
	Dir InstructionType = "dir"
)

type FileEntity struct {
	Name   string
	Type   FileEntityType
	Weight int
	Childs []*FileEntity
}

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day7.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	tree, _ := getTreeFromInstr(&FileEntity{
		Name:   "/",
		Type:   Root,
		Weight: 0,
		Childs: []*FileEntity{},
	}, lines[1:])

	spaceToClean := 30000000 - (70000000 - tree.Weight)

	fmt.Print(getClosestHigher(tree, &FileEntity{
		Name:   "beg",
		Type:   DirType,
		Weight: 10000000000,
		Childs: []*FileEntity{},
	}, spaceToClean))
}

func getTreeFromInstr(root *FileEntity, instrs []string) (tempRoot *FileEntity, instrUsed int) {
	for i := 0; i < len(instrs); i++ {
		instr := instrs[i]

		if strings.HasPrefix(instr, fmt.Sprintf("$ %v", Ls)) {
			continue
		}

		if strings.HasPrefix(instr, fmt.Sprintf("$ %v", Cd)) {
			fileNameForRoot := getFileNameFromCd(instr)

			if fileNameForRoot == ".." {
				return root, i
			}

			rootBelow := findFileByName(root, fileNameForRoot)
			if rootBelow == nil {
				panic("file not found!")
			}

			root.Weight -= rootBelow.Weight
			rootBelow, used := getTreeFromInstr(rootBelow, instrs[i+1:])
			root.Weight += rootBelow.Weight

			if used == -1 {
				return root, -1
			}

			i += used + 1

			continue
		}

		if strings.HasPrefix(instr, fmt.Sprintf("%v ", Dir)) {
			fileName := getFileNameFromDir(instr)

			root.Childs = append(root.Childs, &FileEntity{
				Name:   fileName,
				Weight: 0,
				Type:   DirType,
				Childs: []*FileEntity{},
			})

			continue
		}

		fileName, size := getFileNameWithWeight(instr)
		root.Childs = append(root.Childs, &FileEntity{
			Name:   fileName,
			Weight: size,
			Type:   File,
			Childs: nil,
		})

		root.Weight += size
	}

	return root, -1
}

func sumSpecificSizes(root *FileEntity) int {
	sum := 0

	if root.Weight <= 100000 {
		fmt.Printf("Test: %v:%v \n", root.Name, root.Weight)
		sum += root.Weight
	}

	for _, leaf := range root.Childs {
		if leaf.Type == File {
			continue
		}
		sum += sumSpecificSizes(leaf)
	}

	return sum
}

func getClosestHigher(root *FileEntity, closest *FileEntity, value int) *FileEntity {
	tmpClosest := closest

	if root.Weight >= value && root.Weight < tmpClosest.Weight && root.Type != File {
		tmpClosest = root
	}

	for _, leaf := range root.Childs {
		if leaf.Type == File {
			continue
		}

		tmpClosest = getClosestHigher(leaf, tmpClosest, value)
	}

	return tmpClosest
}

func findFileByName(searchedFile *FileEntity, searchedFileName string) *FileEntity {
	for _, file := range searchedFile.Childs {
		if file.Name == searchedFileName {
			return file
		}
	}

	return nil
}

func getFileNameFromCd(instr string) string {
	return instr[5:]
}

func getFileNameFromDir(instr string) string {
	return instr[4:]
}

func getFileNameWithWeight(instr string) (string, int) {
	size, name, numberEnd := "", "", false

	for _, r := range instr {
		if string(r) == " " {
			numberEnd = true
		}

		if int(r) >= '0' && int(r) <= '9' && !numberEnd {
			size += string(r)
			continue
		}

		name += string(r)
	}

	sizeN, err := strconv.Atoi(size)
	if err != nil {
		panic(err)
	}

	return name, sizeN
}
