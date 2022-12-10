package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ROCK A: 1 Paper B: 2 Scissors C: 3

var (
	figurePoints = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	figuresComb = map[string]int{
		"A X": 3,
		"A Y": 1,
		"A Z": 2,

		"B X": 1,
		"B Y": 2,
		"B Z": 3,

		"C X": 2,
		"C Y": 3,
		"C Z": 1,
	}
)

func main() {
	fmt.Print("Advent of Code****")

	file, err := os.Open("day2.txt")
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

	points := 0
	for _, pairString := range lines {
		figures := strings.Split(pairString, " ")

		points += figurePoints[figures[1]]
		points += figuresComb[pairString]
	}

	fmt.Printf("Result: %v", points)
}
