package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day10.txt")
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

	drawCRT(lines)

}

func countStrengthSignal(instrs []string) int {
	strSig := 0
	xReg := 1
	addxTodo := []int{}

	instrIndex := 0
	for cycle := 0; instrIndex < len(instrs); cycle++ {
		if isReadCycle(cycle) {
			strSig += cycle * xReg
		}

		if len(addxTodo) > 0 {
			val := addxTodo[0]
			addxTodo = addxTodo[1:]
			xReg += val
			instrIndex++
			continue
		}

		in, val := readInstr(instrs[instrIndex])

		if in == "noop" {
			instrIndex++
			continue
		}

		addxTodo = append(addxTodo, val)
	}

	return strSig
}

func drawCRT(instrs []string) {
	xReg := 1
	addxTodo := []int{}
	crt := 0

	instrIndex := 0
	for cycle := 0; instrIndex < len(instrs); cycle++ {
		if cycle%40 == 0 {
			crt = 0
			fmt.Print("\n")
		}

		// fmt.Printf("%v:%v:", cycle, xReg)
		if crt-1 == xReg || crt == xReg || crt+1 == xReg {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if len(addxTodo) > 0 {
			val := addxTodo[0]
			addxTodo = addxTodo[1:]
			xReg += val
			instrIndex++
			crt++
			continue
		}

		in, val := readInstr(instrs[instrIndex])

		if in == "noop" {
			instrIndex++
			crt++
			continue
		}

		crt++
		addxTodo = append(addxTodo, val)
	}
}

func isReadCycle(cycle int) bool {
	return cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220
}

func readInstr(instr string) (string, int) {
	inArr := strings.Split(instr, " ")

	if inArr[0] == "noop" {
		return inArr[0], 0
	}

	inVal, err := strconv.Atoi(inArr[1])
	if err != nil {
		panic(err)
	}

	return inArr[0], inVal
}
