package day11

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

// Chinease theorem
var (
	nums = []int{2, 3, 5, 7, 11, 13, 17, 19}
	prod = getProd()
	pp   = getPP()
	inv  = getNaiveInv()
)

func getProd() int {
	result := 1
	for _, num := range nums {
		result = result * num
	}
	return result
}

func getPP() []int {
	pp := []int{}
	for _, num := range nums {
		pp = append(pp, prod/num)
	}

	return pp
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
	for round := 0; round < 10000; round++ {
		for turn := 0; turn < 8; turn++ {
			for _, itemWorryLevel := range monkeys[turn].Items {
				monkeys[turn].InspectionCount++
				worryLevel := monkeys[turn].Operation(itemWorryLevel)
				worryLevel = worryLevel % prod
				newMonkey := monkeys[turn].ChooseMonkey(worryLevel)
				monkeys[turn].Items = monkeys[turn].Items[1:]
				monkeys[newMonkey].Items = append(monkeys[newMonkey].Items, worryLevel)
			}
		}
	}

	printMonkeys()

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

func naiveChineaseTheorem(worryLevel int) int {
	var rem []int

	for _, num := range nums {
		rem = append(rem, worryLevel%num)
	}

	result := 1
	for {
		flag := true

		for i := 0; i < len(nums); i++ {
			if result%nums[i] != rem[i] {
				flag = false
				break
			}
		}

		if flag {
			return result
		}
		result++
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func getNaiveInv() []int {
	inv := []int{}
	for i, num := range nums {
		t := pp[i] % num

		result := 2
		for {
			if (result*pp[i])%num == t {
				inv = append(inv, result)
				break
			}
			result += 1
		}
	}
	return inv
}

func inverseChineaseTheorem(worryLevel int) int {
	var rem []int

	for _, num := range nums {
		rem = append(rem, worryLevel%num)
	}

	t := 0
	for i, r := range rem {
		t += r * pp[i] * inv[i]
	}

	return t % prod
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

func test() int {
	tnum := []int{3, 4, 5}
	trem := []int{2, 3, 1}
	tprod := 60

	tpp := []int{}
	for _, num := range tnum {
		tpp = append(tpp, tprod/num)
	}

	tinv := []int{}
	for i, num := range tnum {
		t := tpp[i] % num

		result := 2
		for {
			if (result*tpp[i])%num == t {
				inv = append(tinv, result)
				break
			}
			result += 1
		}
	}

	t := 0
	for i, r := range trem {
		t += r * tpp[i] * tinv[i]
	}

	return t % tprod
}
