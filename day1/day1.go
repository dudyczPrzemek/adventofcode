package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Advent of Code****")

	file, err := os.Open("day1.txt")
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

	max := [3]int{0, 0, 0}
	tmpMax := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if index := getIndexForSwap(max, tmpMax); index != -1 {
				max[index] = tmpMax
			}

			tmpMax = 0
			continue
		}

		linesValue, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(err)
		}

		tmpMax += linesValue
	}

	fmt.Printf("Result: %v", sumArr(max))
}

func getIndexForSwap(maxArr [3]int, currValue int) int {
	i, min := getMinWithIndex(maxArr)

	if min < currValue {
		return i
	}

	return -1
}

func getMinWithIndex(arr [3]int) (int, int) {
	min := arr[0]
	index := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			index = i
		}
	}

	return index, min
}

func sumArr(arr [3]int) int {
	sum := 0
	for _, val := range arr {
		fmt.Printf("%v \n", val)
		sum += val
	}
	return sum
}
