package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertex struct {
	Height int
	X      int
	Y      int
}

type VertexCalculation struct {
	PreviousVertex   *Vertex
	ShortestDistance int
}

func main() {
	fmt.Print("****Advent of Code**** \n")

	file, err := os.Open("day12.txt")
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

	posX := 0
	posY := 0

	goalX := 0
	goalY := 0
	// graph := make([][]int, len(lines))
	// for r, row := range lines {
	// 	graph[r] = make([]int, len(row))
	// 	for c, col := range row {
	// 		if col == 'S' {
	// 			posX = r
	// 			posY = c
	// 			graph[r][c] = 1
	// 		} else if col == 'E' {
	// 			graph[r][c] = 27
	// 		} else {
	// 			graph[r][c] = int(col) - 96
	// 		}
	// 	}
	// }

	dijkstraArray := map[*Vertex]*VertexCalculation{}
	graph := make([][]*Vertex, len(lines))
	for r, row := range lines {
		graph[r] = make([]*Vertex, len(row))
		for c, col := range row {
			if col == 'S' {
				posX = r
				posY = c
				graph[r][c] = &Vertex{
					Height: 1,
					X:      r,
					Y:      c,
				}

				dijkstraArray[graph[r][c]] = &VertexCalculation{
					PreviousVertex:   nil,
					ShortestDistance: 0,
				}

			} else if col == 'E' {
				graph[r][c] = &Vertex{
					Height: 27,
					X:      r,
					Y:      c,
				}

				goalX = r
				goalY = c

				dijkstraArray[graph[r][c]] = &VertexCalculation{
					PreviousVertex:   nil,
					ShortestDistance: 10000000000000,
				}
			} else {
				graph[r][c] = &Vertex{
					Height: int(col) - 96,
					X:      r,
					Y:      c,
				}

				dijkstraArray[graph[r][c]] = &VertexCalculation{
					PreviousVertex:   nil,
					ShortestDistance: 10000000000000,
				}
			}
		}
	}

	// printGraph(graph)
	fillDijkstraArray(graph, dijkstraArray, posX, posY)

	shortest, previous := printShortest(dijkstraArray, goalX, goalY)

	fmt.Printf("Result :%v:%v", previous, shortest)
}

func printGraph(graph [][]*Vertex) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			fmt.Printf("%v ", graph[i][j].Height)
		}
		fmt.Println()
	}
}

func printShortest(dijkstraArray map[*Vertex]*VertexCalculation, goalX int, goalY int) (int, *Vertex) {
	for v, k := range dijkstraArray {
		if isNext(v, goalX, goalY) && k.PreviousVertex != nil {
			return k.ShortestDistance, k.PreviousVertex
		}
	}

	return -1, nil
}

func isNext(v *Vertex, goalX int, goalY int) bool {
	return (v.X == goalX && v.Y+1 == goalY) ||
		(v.X-1 == goalX && v.Y == goalY) ||
		(v.X == goalX && v.Y-1 == goalY) ||
		(v.X+1 == goalX && v.Y == goalY)
}

type LocalizationTuple struct {
	X int
	Y int
}

func fillDijkstraArray(graph [][]*Vertex, dijkstraArray map[*Vertex]*VertexCalculation, currX int, currY int) bool {
	// Init
	toVisit := []*LocalizationTuple{{X: currX, Y: currY}}
	for len(toVisit) > 0 {
		currentTuple := toVisit[0]
		currentShortestDistance := dijkstraArray[graph[currentTuple.X][currentTuple.Y]].ShortestDistance

		toVisit = mark(graph, dijkstraArray, currentTuple, currentShortestDistance, toVisit)

		// After mark
		toVisit = toVisit[1:]
	}

	return true
}

func mark(graph [][]*Vertex, dijkstraArray map[*Vertex]*VertexCalculation, currentLoc *LocalizationTuple, currentShortestDistance int, toVisit []*LocalizationTuple) []*LocalizationTuple {
	if currentLoc.Y-1 >= 0 &&
		isSameOrAbove(graph[currentLoc.X][currentLoc.Y], graph[currentLoc.X][currentLoc.Y-1]) &&
		dijkstraArray[graph[currentLoc.X][currentLoc.Y-1]].ShortestDistance > currentShortestDistance+1 {
		dijkstraArray[graph[currentLoc.X][currentLoc.Y-1]].ShortestDistance = currentShortestDistance + 1
		dijkstraArray[graph[currentLoc.X][currentLoc.Y-1]].PreviousVertex = graph[currentLoc.X][currentLoc.Y]
		// fmt.Printf("%v, %v::%v \n", currentLoc.X, currentLoc.Y-1, dijkstraArray[graph[currentLoc.X][currentLoc.Y-1]])
		toVisit = append(toVisit, &LocalizationTuple{X: currentLoc.X, Y: currentLoc.Y - 1})
	}

	if currentLoc.Y+1 < len(graph[currentLoc.X]) &&
		isSameOrAbove(graph[currentLoc.X][currentLoc.Y], graph[currentLoc.X][currentLoc.Y+1]) &&
		dijkstraArray[graph[currentLoc.X][currentLoc.Y+1]].ShortestDistance > currentShortestDistance+1 {
		dijkstraArray[graph[currentLoc.X][currentLoc.Y+1]].ShortestDistance = currentShortestDistance + 1
		dijkstraArray[graph[currentLoc.X][currentLoc.Y+1]].PreviousVertex = graph[currentLoc.X][currentLoc.Y]
		// fmt.Printf("%v, %v::%v \n", currentLoc.X, currentLoc.Y+1, dijkstraArray[graph[currentLoc.X][currentLoc.Y+1]])
		toVisit = append(toVisit, &LocalizationTuple{X: currentLoc.X, Y: currentLoc.Y + 1})
	}

	if currentLoc.X+1 < len(graph) &&
		isSameOrAbove(graph[currentLoc.X][currentLoc.Y], graph[currentLoc.X+1][currentLoc.Y]) &&
		dijkstraArray[graph[currentLoc.X+1][currentLoc.Y]].ShortestDistance > currentShortestDistance+1 {
		dijkstraArray[graph[currentLoc.X+1][currentLoc.Y]].ShortestDistance = currentShortestDistance + 1
		dijkstraArray[graph[currentLoc.X+1][currentLoc.Y]].PreviousVertex = graph[currentLoc.X][currentLoc.Y]
		// fmt.Printf("%v, %v::%v \n", currentLoc.X+1, currentLoc.Y, dijkstraArray[graph[currentLoc.X+1][currentLoc.Y]])
		toVisit = append(toVisit, &LocalizationTuple{X: currentLoc.X + 1, Y: currentLoc.Y})
	}

	if currentLoc.X-1 >= 0 &&
		isSameOrAbove(graph[currentLoc.X][currentLoc.Y], graph[currentLoc.X-1][currentLoc.Y]) &&
		dijkstraArray[graph[currentLoc.X-1][currentLoc.Y]].ShortestDistance > currentShortestDistance+1 {
		dijkstraArray[graph[currentLoc.X-1][currentLoc.Y]].ShortestDistance = currentShortestDistance + 1
		dijkstraArray[graph[currentLoc.X-1][currentLoc.Y]].PreviousVertex = graph[currentLoc.X][currentLoc.Y]
		// fmt.Printf("%v, %v::%v \n", currentLoc.X-1, currentLoc.Y, dijkstraArray[graph[currentLoc.X-1][currentLoc.Y]])
		toVisit = append(toVisit, &LocalizationTuple{X: currentLoc.X - 1, Y: currentLoc.Y})
	}

	return toVisit
}

func isSameOrAbove(v1 *Vertex, v2 *Vertex) bool {
	return v1.Height == v2.Height || v1.Height+1 == v2.Height
}
