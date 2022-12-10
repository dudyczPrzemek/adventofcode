package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	Row   int
	Col   int
	Score int
}

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day8.txt")
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

	grid := transformInput(lines)

	visibleTrees, _ := countVisibleTrees(grid)

	result := countScenicScore(grid, visibleTrees)
	fmt.Printf("Result: %v", result)
}

func transformInput(input []string) [][]int {
	result := make([][]int, len(input))
	for i := range result {
		result[i] = make([]int, len(input[0]))
	}

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			val, err := strconv.Atoi(fmt.Sprintf("%c", input[row][col]))
			if err != nil {
				panic(err)
			}

			result[row][col] = val
		}
	}

	return result
}

func countVisibleTrees(grid [][]int) ([]*Tree, int) {
	sum := 4*len(grid) - 4

	visibleTrees := []*Tree{}

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			val := grid[row][col]

			left := findHighest(grid[row][:col])
			if left < val {
				sum += 1
				visibleTrees = append(visibleTrees, &Tree{Row: row, Col: col, Score: 0})
				continue
			}

			right := findHighest(grid[row][col+1:])
			if right < val {
				sum += 1
				visibleTrees = append(visibleTrees, &Tree{Row: row, Col: col, Score: 0})
				continue
			}

			top := findHighestTop(grid, row-1, col)
			if top < val {
				sum += 1
				visibleTrees = append(visibleTrees, &Tree{Row: row, Col: col, Score: 0})
				continue
			}

			bottom := findHighestBottom(grid, row+1, col)
			if bottom < val {
				sum += 1
				visibleTrees = append(visibleTrees, &Tree{Row: row, Col: col, Score: 0})
				continue
			}

		}
	}

	return visibleTrees, sum
}

func countScenicScore(grid [][]int, visibleTrees []*Tree) int {
	highestScore := 0

	for _, tree := range visibleTrees {
		val := grid[tree.Row][tree.Col]

		left := countScenicRowLeft(grid[tree.Row][:tree.Col], val)

		right := countScenicRowRight(grid[tree.Row][tree.Col+1:], val)

		top := countScenicTop(grid, tree.Row-1, tree.Col, val)

		bottom := countScenicBottom(grid, tree.Row+1, tree.Col, val)

		currScore := left * right * top * bottom

		if highestScore < currScore {
			highestScore = currScore
		}
	}

	return highestScore
}

func countScenicRowRight(arr []int, currHeight int) int {
	score := 0
	for _, val := range arr {
		if val >= currHeight {
			score += 1
			return score
		}

		score += 1
	}

	return score
}

func countScenicRowLeft(arr []int, currHeight int) int {
	score := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= currHeight {
			score += 1
			return score
		}

		score += 1
	}

	return score
}

func findHighest(arr []int) int {
	highest := 0
	for _, val := range arr {
		if val > highest {
			highest = val
		}
	}

	return highest
}

func findHighestTop(grid [][]int, startRow int, col int) int {
	highest := 0
	for i := startRow; i >= 0; i-- {
		if grid[i][col] > highest {
			highest = grid[i][col]
		}
	}

	return highest
}

func countScenicTop(grid [][]int, startRow int, col int, currHeight int) int {
	score := 0
	for i := startRow; i >= 0; i-- {
		if grid[i][col] >= currHeight {
			score += 1
			return score
		}
		score += 1
	}

	return score
}

func findHighestBottom(grid [][]int, startRow int, col int) int {
	highest := 0
	for i := startRow; i < len(grid); i++ {
		if grid[i][col] > highest {
			highest = grid[i][col]
		}
	}

	return highest
}

func countScenicBottom(grid [][]int, startRow int, col int, currHeight int) int {
	score := 0
	for i := startRow; i < len(grid); i++ {
		if grid[i][col] >= currHeight {
			score += 1
			return score
		}
		score += 1
	}

	return score
}
