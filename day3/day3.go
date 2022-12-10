package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day3.txt")
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
	// for _, rucksack := range lines {
	// 	duplicateChar := getDuplicates(rucksack)

	// 	if duplicateChar != "" {
	// 		asciiVal := []byte(duplicateChar)[0]

	// 		switch {
	// 		case asciiVal >= 97 && asciiVal <= 122:
	// 			sum += int(asciiVal) - 96
	// 			break
	// 		case asciiVal >= 65 && asciiVal <= 90:
	// 			sum += int(asciiVal) - 38
	// 			break
	// 		default:
	// 			log.Fatalf("Invalid character: %v", duplicateChar)
	// 		}
	// 	}
	// }

	for i := 0; i < len(lines); i = i + 3 {
		rucksacks := lines[i : i+3]

		badge := getBadge(rucksacks)

		if badge != "" {
			asciiVal := []byte(badge)[0]

			switch {
			case asciiVal >= 97 && asciiVal <= 122:
				sum += int(asciiVal) - 96
				break
			case asciiVal >= 65 && asciiVal <= 90:
				sum += int(asciiVal) - 38
				break
			default:
				log.Fatalf("Invalid character: %v", badge)
			}
		}
	}

	fmt.Printf("Result: %v", sum)

}

func getDuplicates(rucksack string) string {
	compLen := len(rucksack) / 2
	compareMapfFirst := map[string]bool{}
	compareMapfSecond := map[string]bool{}

	for i := 0; i < compLen; i++ {
		first := fmt.Sprintf("%c ", rucksack[i])
		second := fmt.Sprintf("%c ", rucksack[i+compLen])

		if compareMapfSecond[first] || first == second {
			return first
		}

		if compareMapfFirst[second] {
			return second
		}

		compareMapfFirst[first] = true
		compareMapfSecond[second] = true
	}

	return ""
}

func getBadge(rucksacks []string) string {
	compareMap := map[string]int{}

	for _, rucksack := range rucksacks {
		tempRucksack := map[string]bool{}

		for i := 0; i < len(rucksack); i++ {
			char := fmt.Sprintf("%c ", rucksack[i])

			if !tempRucksack[char] {
				compareMap[char] = compareMap[char] + 1

				if compareMap[char] == 3 {
					return char
				}

				tempRucksack[char] = true
			}
		}
	}

	return ""
}
