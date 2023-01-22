package main

import (
	"bufio"
	"fmt"
	"os"
)

type Monkey struct {
	Items           []int
	InspectionCount int
	Operation       func(int) int
	ChooseMonkey    func(int) int
}

var monkeys = [8]*Monkey{
	{
		Items:           []int{66, 79},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old * 11
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%7 == 0 {
				return 6
			}
			return 7
		},
	},
	{
		Items:           []int{84, 94, 94, 81, 98, 75},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old * 17
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%13 == 0 {
				return 5
			}
			return 2
		},
	},
	{
		Items:           []int{85, 79, 59, 64, 79, 95, 67},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old + 8
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%5 == 0 {
				return 4
			}
			return 5
		},
	},
	{
		Items:           []int{70},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old + 3
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%19 == 0 {
				return 6
			}
			return 0
		},
	},
	{
		Items:           []int{57, 69, 78, 78},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old + 4
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%2 == 0 {
				return 0
			}
			return 3
		},
	},
	{
		Items:           []int{65, 92, 60, 74, 72},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old + 7
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%11 == 0 {
				return 3
			}
			return 4
		},
	},
	{
		Items:           []int{77, 91, 91},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old * old
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%17 == 0 {
				return 1
			}
			return 7
		},
	},
	{
		Items:           []int{76, 58, 57, 55, 67, 77, 54, 99},
		InspectionCount: 0,
		Operation: func(old int) int {
			return old + 6
		},
		ChooseMonkey: func(worryLevel int) int {
			if worryLevel%3 == 0 {
				return 2
			}
			return 1
		},
	},
}

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day11.txt")
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

	result := getMonkeyBusiness()

	fmt.Printf("Result: %v", result)
}

func getMonkeyBusiness() int {
	for round := 0; round < 20; round++ {
		for turn := 0; turn < 8; turn++ {
			for _, itemWorryLevel := range monkeys[turn].Items {
				monkeys[turn].InspectionCount++
				worryLevel := monkeys[turn].Operation(itemWorryLevel)
				worryLevel = worryLevel / 3
				newMonkey := monkeys[turn].ChooseMonkey(worryLevel)
				monkeys[turn].Items = monkeys[turn].Items[1:]
				monkeys[newMonkey].Items = append(monkeys[newMonkey].Items, worryLevel)
			}
		}

		printMonkeys()

	}

	mostActives := getTwoMostActive(monkeys)

	return monkeys[mostActives[0]].InspectionCount * monkeys[mostActives[1]].InspectionCount
}

func getTwoMostActive(monkeys [8]*Monkey) [2]int {
	mostActives := [2]int{0, 0}

	for index, monkey := range monkeys {
		if monkeys[mostActives[0]].InspectionCount < monkey.InspectionCount {
			mostActives[0] = index
			continue
		}

		if monkeys[mostActives[1]].InspectionCount < monkey.InspectionCount {
			mostActives[1] = index
		}
	}

	return mostActives
}

func printMonkeys() {
	fmt.Printf("%v\n", "-----------------------------------")
	for _, monkey := range monkeys {
		for _, itemWorryLevel := range monkey.Items {
			fmt.Printf("%v, ", itemWorryLevel)
		}
		fmt.Printf("::%v\n", monkey.InspectionCount)
	}

	fmt.Printf("%v\n", "-----------------------------------")
}
