package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GridCoordinate struct {
	y int
	x int
}

type Rope struct {
	knots []*GridCoordinate
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func debugGrid(rope *Rope) {
	var grid [21][26]string

	for i := 0; i < 21; i++ {
		for j := 0; j < 26; j++ {
			grid[i][j] = "."
		}
	}

	for i, knot := range rope.knots {
		grid[knot.y][knot.x] = strconv.Itoa(9 - i)
	}

	for index, row := range grid {
		fmt.Println(row, index)
	}
}

func isCoordinateNear(gc1 *GridCoordinate, gc2 *GridCoordinate) bool {
	if absDiffInt(gc1.x, gc2.x) < 2 && absDiffInt(gc1.y, gc2.y) < 2 {
		return true
	}

	return false
}

func (rope *Rope) moveLR(currentKnotIndex int, dir int) {
	if currentKnotIndex == -1 {
		return
	}

	hIndex := len(rope.knots) - 1

	if hIndex == currentKnotIndex {
		rope.knots[hIndex].x = rope.knots[hIndex].x + dir
		rope.moveLR(currentKnotIndex-1, dir)
		return
	}

	t := rope.knots[currentKnotIndex]
	h := rope.knots[currentKnotIndex+1]

	if isCoordinateNear(t, h) {
		rope.moveLR(currentKnotIndex-1, dir)
		return
	}

	if t.x < h.x {
		t.x++
	}

	if t.x > h.x {
		t.x--
	}

	if h.y < t.y {
		t.y--
	}

	if h.y > t.y {
		t.y++
	}

	rope.moveLR(currentKnotIndex-1, dir)
}

func (rope *Rope) moveUD(currentKnotIndex int, dir int) {
	if currentKnotIndex == -1 {
		return
	}

	hIndex := len(rope.knots) - 1

	if hIndex == currentKnotIndex {
		rope.knots[hIndex].y = rope.knots[hIndex].y + dir
		rope.moveUD(currentKnotIndex-1, dir)
		return
	}

	t := rope.knots[currentKnotIndex]
	h := rope.knots[currentKnotIndex+1]

	if isCoordinateNear(t, h) {
		rope.moveUD(currentKnotIndex-1, dir)
		return
	}

	if t.x < h.x {
		t.x++
	}

	if t.x > h.x {
		t.x--
	}

	if h.y < t.y {
		t.y--
	}

	if h.y > t.y {
		t.y++
	}

	rope.moveUD(currentKnotIndex-1, dir)
}

func move(rope *Rope, howMany int, direction string) []GridCoordinate {
	tailCoordinates := []GridCoordinate{}
	hKnotIndex := len(rope.knots) - 1
	tKnotIndex := 0

	for i := 0; i < howMany; i++ {
		if direction == "L" {
			rope.moveLR(hKnotIndex, -1)
		}

		if direction == "R" {
			rope.moveLR(hKnotIndex, 1)
		}

		if direction == "U" {
			rope.moveUD(hKnotIndex, -1)
		}

		if direction == "D" {
			rope.moveUD(hKnotIndex, 1)
		}

		//debugGrid(rope)
		tKnot := rope.knots[tKnotIndex]
		tailCoordinates = append(tailCoordinates, GridCoordinate{y: tKnot.y, x: tKnot.x})
	}

	return tailCoordinates
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	knots := []*GridCoordinate{}

	for i := 0; i < 10; i++ {
		knots = append(knots, &GridCoordinate{y: 15, x: 11})
	}

	rope := Rope{knots: knots}
	allMoves := []GridCoordinate{}

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		splittedLine := strings.Fields(line)
		direction := splittedLine[0]
		numberOfSteps, _ := strconv.Atoi(splittedLine[1])
		moveList := move(&rope, numberOfSteps, direction)

		allMoves = append(allMoves, moveList...)
	}

	uniqueMoves := make(map[string]struct{})

	for _, move := range allMoves {
		k := fmt.Sprintf("%d - %d", move.y, move.x)
		uniqueMoves[k] = struct{}{}
	}

	fmt.Println(len(uniqueMoves))

	readFile.Close()

}
