package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instr struct {
	Dir   string
	Steps int
}

type Field struct {
	X int
	Y int
}

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day9.txt")
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

	instrs := transformInstr(lines)

	fmt.Printf("Result: %v", get9TailPositions(instrs))

}

func transformInstr(inputs []string) []*Instr {
	result := []*Instr{}

	for _, input := range inputs {
		inArr := strings.Split(input, " ")
		steps, err := strconv.Atoi(inArr[1])
		if err != nil {
			panic(err)
		}

		result = append(result, &Instr{
			Dir:   inArr[0],
			Steps: steps,
		})
	}

	return result
}

func getTailPositions(instrs []*Instr) int {
	hPos := &Field{
		X: 0,
		Y: 0,
	}

	tPos := &Field{
		X: 0,
		Y: 0,
	}

	tVis := map[string]bool{
		"0:0": true,
	}

	for _, instr := range instrs {
		switch {
		case instr.Dir == "U":
			for i := 0; i < instr.Steps; i++ {
				hPos.Y += 1
				tPos = adjustTailPosition(hPos, tPos)

				tVis[fmt.Sprintf("%v:%v", tPos.X, tPos.Y)] = true
			}
		case instr.Dir == "R":
			for i := 0; i < instr.Steps; i++ {
				hPos.X += 1
				tPos = adjustTailPosition(hPos, tPos)

				tVis[fmt.Sprintf("%v:%v", tPos.X, tPos.Y)] = true
			}
		case instr.Dir == "L":
			for i := 0; i < instr.Steps; i++ {
				hPos.X -= 1
				tPos = adjustTailPosition(hPos, tPos)

				tVis[fmt.Sprintf("%v:%v", tPos.X, tPos.Y)] = true
			}
		case instr.Dir == "D":
			for i := 0; i < instr.Steps; i++ {
				hPos.Y -= 1
				tPos = adjustTailPosition(hPos, tPos)

				tVis[fmt.Sprintf("%v:%v", tPos.X, tPos.Y)] = true
			}
		}
	}

	return len(tVis)
}

func get9TailPositions(instrs []*Instr) int {
	hPos := &Field{
		X: 0,
		Y: 0,
	}

	arrTPost := make([]*Field, 9)
	for i := 0; i < 9; i++ {
		arrTPost[i] = &Field{
			X: 0,
			Y: 0,
		}
	}

	tVis := map[string]bool{
		"0:0": true,
	}

	for _, instr := range instrs {
		switch {
		case instr.Dir == "U":
			for i := 0; i < instr.Steps; i++ {
				hPos.Y += 1
				arrTPost = adjustTailsPosition(hPos, arrTPost)

				tVis[fmt.Sprintf("%v:%v", arrTPost[8].X, arrTPost[8].Y)] = true
			}
		case instr.Dir == "R":
			for i := 0; i < instr.Steps; i++ {
				hPos.X += 1
				arrTPost = adjustTailsPosition(hPos, arrTPost)

				tVis[fmt.Sprintf("%v:%v", arrTPost[8].X, arrTPost[8].Y)] = true
			}
		case instr.Dir == "L":
			for i := 0; i < instr.Steps; i++ {
				hPos.X -= 1
				arrTPost = adjustTailsPosition(hPos, arrTPost)

				tVis[fmt.Sprintf("%v:%v", arrTPost[8].X, arrTPost[8].Y)] = true
			}
		case instr.Dir == "D":
			for i := 0; i < instr.Steps; i++ {
				hPos.Y -= 1
				arrTPost = adjustTailsPosition(hPos, arrTPost)

				tVis[fmt.Sprintf("%v:%v", arrTPost[8].X, arrTPost[8].Y)] = true
			}
		}
	}

	return len(tVis)
}

func adjustTailsPosition(hPos *Field, tArr []*Field) []*Field {
	for i := 0; i < len(tArr); i++ {
		if i == 0 {
			tArr[i] = adjustTailPosition(hPos, tArr[i])
			continue
		}

		tArr[i] = adjustTailPosition(tArr[i-1], tArr[i])
	}

	return tArr
}

func recognizeDir(hPos *Field, tPos *Field) string {
	if hPos.X == tPos.X+2 && hPos.Y == tPos.Y+1 ||
		hPos.X == tPos.X+2 && hPos.Y == tPos.Y-1 ||
		hPos.X == tPos.X+2 && hPos.Y == tPos.Y {
		return "R"
	}

	if hPos.X == tPos.X-2 && hPos.Y == tPos.Y+1 ||
		hPos.X == tPos.X-2 && hPos.Y == tPos.Y-1 ||
		hPos.X == tPos.X-2 && hPos.Y == tPos.Y {
		return "L"
	}

	if hPos.X == tPos.X+1 && hPos.Y == tPos.Y+2 ||
		hPos.X == tPos.X-1 && hPos.Y == tPos.Y+2 ||
		hPos.X == tPos.X && hPos.Y == tPos.Y+2 {
		return "U"
	}

	if hPos.X == tPos.X+1 && hPos.Y == tPos.Y-2 ||
		hPos.X == tPos.X-1 && hPos.Y == tPos.Y-2 ||
		hPos.X == tPos.X && hPos.Y == tPos.Y-2 {
		return "D"
	}

	if hPos.X == tPos.X+2 && hPos.Y == tPos.Y+2 {
		return "UR"
	}

	if hPos.X == tPos.X-2 && hPos.Y == tPos.Y+2 {
		return "UL"
	}

	if hPos.X == tPos.X+2 && hPos.Y == tPos.Y-2 {
		return "DR"
	}

	if hPos.X == tPos.X-2 && hPos.Y == tPos.Y-2 {
		return "DL"
	}

	return ""
}

func adjustTailPosition(hPos *Field, tPos *Field) *Field {
	//Same field
	if hPos.X == tPos.X && hPos.Y == tPos.Y {
		return tPos
	}

	//Next to each other vertically
	if hPos.X == tPos.X && (hPos.Y == tPos.Y+1 || hPos.Y == tPos.Y-1) {
		return tPos
	}

	//Next to each other horizotally
	if hPos.Y == tPos.Y && (hPos.X == tPos.X+1 || hPos.X == tPos.X-1) {
		return tPos
	}

	//Next to each other diagonnaly
	if (hPos.X == tPos.X+1 && hPos.Y == tPos.Y+1) ||
		(hPos.X == tPos.X-1 && hPos.Y == tPos.Y+1) ||
		(hPos.X == tPos.X-1 && hPos.Y == tPos.Y-1) ||
		(hPos.X == tPos.X+1 && hPos.Y == tPos.Y-1) {
		return tPos
	}

	switch recognizeDir(hPos, tPos) {
	case "U":
		tPos.X = hPos.X
		tPos.Y = hPos.Y - 1
	case "D":
		tPos.X = hPos.X
		tPos.Y = hPos.Y + 1
	case "R":
		tPos.X = hPos.X - 1
		tPos.Y = hPos.Y
	case "L":
		tPos.X = hPos.X + 1
		tPos.Y = hPos.Y
	case "UR":
		tPos.X = hPos.X - 1
		tPos.Y = hPos.Y - 1
	case "UL":
		tPos.X = hPos.X + 1
		tPos.Y = hPos.Y - 1
	case "DR":
		tPos.X = hPos.X - 1
		tPos.Y = hPos.Y + 1
	case "DL":
		tPos.X = hPos.X + 1
		tPos.Y = hPos.Y + 1
	}

	return tPos
}
