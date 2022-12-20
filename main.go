package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

type Monkey struct {
	Items           []*big.Int
	InspectionCount int
	Operation       func(*big.Int) *big.Int
	ChooseMonkey    func(*big.Int) int
}

var monkeys = [8]*Monkey{
	{
		Items:           []*big.Int{big.NewInt(66), big.NewInt(79)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Mul(old, big.NewInt(11))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(7)).Cmp(big.NewInt(0)) == 0 {
				return 6
			}
			return 7
		},
	},
	{
		Items:           []*big.Int{big.NewInt(84), big.NewInt(94), big.NewInt(94), big.NewInt(81), big.NewInt(98), big.NewInt(75)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Mul(old, big.NewInt(17))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(13)).Cmp(big.NewInt(0)) == 0 {
				return 5
			}
			return 2
		},
	},
	{
		Items:           []*big.Int{big.NewInt(85), big.NewInt(79), big.NewInt(59), big.NewInt(64), big.NewInt(79), big.NewInt(95), big.NewInt(67)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(8))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(5)).Cmp(big.NewInt(0)) == 0 {
				return 4
			}
			return 5
		},
	},
	{
		Items:           []*big.Int{big.NewInt(70)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(3))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(19)).Cmp(big.NewInt(0)) == 0 {
				return 6
			}
			return 0
		},
	},
	{
		Items:           []*big.Int{big.NewInt(57), big.NewInt(69), big.NewInt(78), big.NewInt(78)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(4))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
				return 0
			}
			return 3
		},
	},
	{
		Items:           []*big.Int{big.NewInt(65), big.NewInt(92), big.NewInt(60), big.NewInt(74), big.NewInt(72)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(7))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(11)).Cmp(big.NewInt(0)) == 0 {
				return 3
			}
			return 4
		},
	},
	{
		Items:           []*big.Int{big.NewInt(77), big.NewInt(91), big.NewInt(91)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Mul(old, old)
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(17)).Cmp(big.NewInt(0)) == 0 {
				return 1
			}
			return 7
		},
	},
	{
		Items:           []*big.Int{big.NewInt(76), big.NewInt(58), big.NewInt(57), big.NewInt(55), big.NewInt(67), big.NewInt(77), big.NewInt(54), big.NewInt(99)},
		InspectionCount: 0,
		Operation: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(6))
		},
		ChooseMonkey: func(worryLevel *big.Int) int {
			if big.NewInt(0).Mod(worryLevel, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
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
	for round := 0; round < 10000; round++ {
		for turn := 0; turn < 8; turn++ {
			for _, itemWorryLevel := range monkeys[turn].Items {
				monkeys[turn].InspectionCount++
				monkeys[turn].Operation(itemWorryLevel)
				itemWorryLevel.Div(itemWorryLevel, big.NewInt(10))
				newMonkey := monkeys[turn].ChooseMonkey(itemWorryLevel)
				monkeys[turn].Items = monkeys[turn].Items[1:]
				monkeys[newMonkey].Items = append(monkeys[newMonkey].Items, itemWorryLevel)
			}
		}
		fmt.Println(round)
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
			fmt.Printf("%v, ", itemWorryLevel.String())
		}
		fmt.Printf("::%v\n", monkey.InspectionCount)
	}

	fmt.Printf("%v\n", "-----------------------------------")
}
