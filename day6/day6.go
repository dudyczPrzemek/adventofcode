package day6

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day6.txt")
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

	result := getPacketStart(lines[0])

	fmt.Printf("Result: %v", result)
}

func getPacketStart(message string) int {
	packetStart := []string{}

	for i := 0; i < len(message); i++ {
		currChar := fmt.Sprintf("%c", message[i])

		if eli := indexOf(currChar, packetStart); eli != -1 {
			packetStart = packetStart[eli+1:]
		}

		if indexOf(currChar, packetStart) == -1 {
			packetStart = append(packetStart, currChar)

			if len(packetStart) == 14 {
				fmt.Print(packetStart)
				return i
			}
		}
	}

	return -1
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
