package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day4.txt")
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

	sum := 0
	for _, ranges := range lines {
		rans := strings.Split(ranges, ",")
		frans := strings.Split(rans[0], "-")
		srans := strings.Split(rans[1], "-")

		frd, err := strconv.Atoi(frans[0])
		if err != nil {
			panic(err)
		}

		fru, err := strconv.Atoi(frans[1])
		if err != nil {
			panic(err)
		}

		srd, err := strconv.Atoi(srans[0])
		if err != nil {
			panic(err)
		}

		sru, err := strconv.Atoi(srans[1])
		if err != nil {
			panic(err)
		}

		if doRangesOverlap([2]int{frd, fru}, [2]int{srd, sru}) || doRangesOverlap([2]int{srd, sru}, [2]int{frd, fru}) {
			sum += 1
		}
	}

	fmt.Printf("Result: %v", sum)
}

func doRangesFullyOverlap(frange [2]int, srange [2]int) bool {
	return frange[0] >= srange[0] && frange[1] <= srange[1]
}

func doRangesOverlap(frange [2]int, srange [2]int) bool {
	return (frange[0] >= srange[0] && frange[0] <= srange[1]) || (frange[1] >= srange[0] && frange[1] <= srange[1])
}
