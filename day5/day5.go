package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var i = 11

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day5.txt")
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

	puzzleEnd := getEndOfPuzzle(lines)
	instrStart := puzzleEnd + 2
	stocks := getStocks(lines[:puzzleEnd])

	getResultStock9001(stocks, lines[instrStart:len(lines)])
	// result := ""
	// for _, rstock := range stocks {
	// 	if len(rstock) == 0 {
	// 		result += " "
	// 	}
	// 	result += rstock[0]
	// }

	// fmt.Printf("Result: %v", result)

	fmt.Print(stocks)
}

func getEndOfPuzzle(inputs []string) int {
	for index, input := range inputs {
		if input == "" {
			return index - 1
		}
	}
	return -1
}

func getStocks(inputs []string) map[int][]string {
	stockLen := len(inputs[0])
	result := map[int][]string{}

	for _, input := range inputs {
		stockIndex := 0
		for i := 1; i < stockLen; i = i + 4 {
			if fmt.Sprintf("%c", input[i]) != " " {
				result[stockIndex] = append(result[stockIndex], fmt.Sprintf("%c", input[i]))
			}

			stockIndex++
		}
	}

	return result
}

func transformInstr(instr string) (numb int, from int, to int) {
	instrValue := [3]string{"", "", ""}

	index := 0
	found := false
	for _, r := range instr {
		if int(r) >= '0' && int(r) <= '9' {
			found = true
			instrValue[index] += string(r)
			continue
		}

		if found {
			index++
			found = false
		}
	}

	numb, err := strconv.Atoi(instrValue[0])
	if err != nil {
		panic(err)
	}

	from, err = strconv.Atoi(instrValue[1])
	if err != nil {
		panic(err)
	}
	to, err = strconv.Atoi(instrValue[2])
	if err != nil {
		panic(err)
	}

	return numb, from - 1, to - 1
}

func getResultStock(begStock map[int][]string, instrs []string) {
	for _, instr := range instrs {
		numb, from, to := transformInstr(instr)

		for i := 0; i < numb; i++ {
			element := begStock[from][0]
			begStock[from] = begStock[from][1:len(begStock[from])]
			begStock[to] = append([]string{element}, begStock[to]...)
		}
	}
}

func getResultStock9001(begStock map[int][]string, instrs []string) {
	for _, instr := range instrs {
		numb, from, to := transformInstr(instr)

		fmt.Printf("%v:%v:%v \n", numb, from, to)

		elements := append([]string{}, begStock[from][0:numb]...)
		begStock[from] = begStock[from][numb:len(begStock[from])]
		begStock[to] = append(elements, begStock[to]...)

		fmt.Printf("%v %v-%v:%v-%v \n", i, from, begStock[from], to, begStock[to])
		// fmt.Printf("%v \n", elements)
		i++
	}
}
