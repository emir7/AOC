package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	i int
	j int
}

type DijkstraEntry struct {
	shortestDistance int
	previousVertex   string
}

func findLetter(grid *[][]string, letter string) (int, int) {
	for i, row := range *grid {
		for j, c := range row {
			if c == letter {
				return i, j
			}
		}
	}

	return -1, -1
}

func findMin(slice *[]int) int {
	min := math.MaxInt

	for _, v := range *slice {
		if v <= min {
			min = v
		}
	}

	return min
}

func printGrid(grid *[][]bool) {
	for _, row := range *grid {
		fmt.Println(row)
	}
}

func getValidMoves(grid *[][]string, visited *[][]bool, cI int, cJ int) []*Move {
	moveR := cJ + 1
	moveL := cJ - 1
	moveU := cI - 1
	moveD := cI + 1
	wid := len((*grid)[0])
	hei := len(*grid)
	moves := []*Move{}

	v1 := int([]rune((*grid)[cI][cJ])[0]) - 97

	if (*grid)[cI][cJ] == "S" {
		v1 = 0
	}

	if moveR < wid && !(*visited)[cI][moveR] {
		v2 := int([]rune((*grid)[cI][moveR])[0]) - 97

		if (*grid)[cI][moveR] == "E" {
			v2 = 26
		}

		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: cI, j: moveR})
		}
	}

	if moveD < hei && !(*visited)[moveD][cJ] {
		v2 := int([]rune((*grid)[moveD][cJ])[0]) - 97

		if (*grid)[moveD][cJ] == "E" {
			v2 = 26
		}

		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: moveD, j: cJ})
		}
	}

	if moveL >= 0 && !(*visited)[cI][moveL] {
		v2 := int([]rune((*grid)[cI][moveL])[0]) - 97
		if (*grid)[cI][moveL] == "E" {
			v2 = 26
		}
		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: cI, j: moveL})
		}
	}

	if moveU >= 0 && !(*visited)[moveU][cJ] {
		v2 := int([]rune((*grid)[moveU][cJ])[0]) - 97
		if (*grid)[moveU][cJ] == "E" {
			v2 = 26
		}
		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: moveU, j: cJ})
		}
	}

	return moves
}

func getValidMoves2(grid *[][]string, cI int, cJ int) []*Move {
	moveR := cJ + 1
	moveL := cJ - 1
	moveU := cI - 1
	moveD := cI + 1
	wid := len((*grid)[0])
	hei := len(*grid)
	moves := []*Move{}

	v1 := int([]rune((*grid)[cI][cJ])[0]) - 97

	if (*grid)[cI][cJ] == "S" {
		v1 = 0
	}

	if moveR < wid {
		v2 := int([]rune((*grid)[cI][moveR])[0]) - 97

		if (*grid)[cI][moveR] == "E" {
			v2 = 26
		}

		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: cI, j: moveR})
		}
	}

	if moveD < hei {
		v2 := int([]rune((*grid)[moveD][cJ])[0]) - 97

		if (*grid)[moveD][cJ] == "E" {
			v2 = 26
		}

		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: moveD, j: cJ})
		}
	}

	if moveL >= 0 {
		v2 := int([]rune((*grid)[cI][moveL])[0]) - 97
		if (*grid)[cI][moveL] == "E" {
			v2 = 26
		}
		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: cI, j: moveL})
		}
	}

	if moveU >= 0 {
		v2 := int([]rune((*grid)[moveU][cJ])[0]) - 97
		if (*grid)[moveU][cJ] == "E" {
			v2 = 26
		}
		if v2-v1 <= 1 {
			moves = append(moves, &Move{i: moveU, j: cJ})
		}
	}

	return moves
}

func findShortestPath1(grid *[][]string, visited *[][]bool, cI int, cJ int, eI int, eJ int, depth int) bool {
	if cI == eI && cJ == eJ {
		return true
	}

	if depth == 0 {
		return false
	}

	validMoves := getValidMoves(grid, visited, cI, cJ)
	for _, move := range validMoves {
		(*visited)[move.i][move.j] = true
		foundPath := findShortestPath1(grid, visited, move.i, move.j, eI, eJ, depth-1)
		(*visited)[move.i][move.j] = false

		if foundPath {
			return true
		}

	}

	return false
}

func printTable(resultTable map[string]*DijkstraEntry) {
	fmt.Println("----")
	for key, value := range resultTable {
		fmt.Println(key, "previousVertex = ", value.previousVertex, "shortestDistance = ", value.shortestDistance)
	}

	fmt.Println("----")
}

func getMinEntryKey(resultTable map[string]*DijkstraEntry, visited map[string]struct{}) string {
	min := math.MaxInt
	minKey := ""

	for key, value := range resultTable {
		_, ok := visited[key]

		if ok {
			continue
		}

		if value.shortestDistance <= min {
			min = value.shortestDistance
			minKey = key
		}
	}

	return minKey
}

func dijkstra(grid *[][]string, cI int, cJ int, eI int, eJ int) int {
	resultTable := make(map[string]*DijkstraEntry)
	unvisited := make(map[string]struct{})
	r1 := strconv.Itoa(cI)
	r2 := strconv.Itoa(cJ)
	startKey := r1 + "," + r2

	for i, row := range *grid {
		y := strconv.Itoa(i)
		for j := range row {
			x := strconv.Itoa(j)
			resultTable[y+","+x] = &DijkstraEntry{shortestDistance: math.MaxInt, previousVertex: ""}
			unvisited[y+","+x] = struct{}{}
		}
	}

	resultTable[startKey] = &DijkstraEntry{shortestDistance: 0, previousVertex: startKey}

	visited := make(map[string]struct{})

	for len(unvisited) > 0 {
		nextNode := getMinEntryKey(resultTable, visited)

		if len(nextNode) == 0 {
			break
		}

		splittedKey := strings.Split(nextNode, ",")
		currentPathValue := resultTable[nextNode].shortestDistance
		cI, _ := strconv.Atoi(splittedKey[0])
		cJ, _ := strconv.Atoi(splittedKey[1])

		validMoves := getValidMoves2(grid, cI, cJ)

		for _, move := range validMoves {
			i := strconv.Itoa(move.i)
			j := strconv.Itoa(move.j)
			currentPath := resultTable[i+","+j].shortestDistance
			calculatedPath := currentPathValue + 1

			if calculatedPath < currentPath {
				resultTable[i+","+j].shortestDistance = calculatedPath
				resultTable[i+","+j].previousVertex = nextNode
			}
		}

		visited[nextNode] = struct{}{}
		delete(unvisited, nextNode)
	}

	f1 := strconv.Itoa(eI)
	f2 := strconv.Itoa(eJ)

	return resultTable[f1+","+f2].shortestDistance
}

func findStartingPoints(grid *[][]string) []*Move {
	startingPoints := []*Move{}

	for i, row := range *grid {
		for j, c := range row {
			if c == "S" || c == "a" {
				startingPoints = append(startingPoints, &Move{i: i, j: j})
			}
		}
	}

	return startingPoints
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	grid := [][]string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splittedInput := strings.Split(line, "")

		grid = append(grid, splittedInput)
	}

	sI, sJ := findLetter(&grid, "S")
	eI, eJ := findLetter(&grid, "E")
	visited := [][]bool{}

	for _, row := range grid {
		visitedRow := []bool{}
		for i := 0; i < len(row); i++ {
			visitedRow = append(visitedRow, false)
		}
		visited = append(visited, visitedRow)
	}
	visited[sI][sJ] = true
	fmt.Println(sI, sJ, eI, eJ)

	//depth := 23

	//fmt.Println(findStartingPoints(&grid))
	//dijkstra(&grid, sI, sJ, eI, eJ)
	/*for true {
		f := findShortestPath(&grid, &visited, sI, sJ, eI, eJ, depth)

		if f {
			fmt.Println(depth)
			break
		}
		fmt.Println(depth)
		depth++
	}*/

	startingPoints := findStartingPoints(&grid)
	min := math.MaxInt
	fmt.Println(len(startingPoints))
	c := 0
	for _, p := range startingPoints {
		d := dijkstra(&grid, p.i, p.j, eI, eJ)
		fmt.Println("calculated", d, "current min", min)
		if d >= 0 && d < min {
			min = d
		}
		c++
		fmt.Println(c, "/", len(startingPoints))
	}

	fmt.Println("result", min)
	readFile.Close()

}
