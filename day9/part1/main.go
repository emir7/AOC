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
	h GridCoordinate
	t GridCoordinate
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func isCoordinateNear(gc1 GridCoordinate, gc2 GridCoordinate) bool {
	if absDiffInt(gc1.x, gc2.x) < 2 && absDiffInt(gc1.y, gc2.y) < 2 {
		return true
	}

	return false
}

func (rope *Rope) moveRight() {
	rope.h.x++

	if isCoordinateNear(rope.t, rope.h) {
		return
	}

	rope.t.x++

	if rope.h.y < rope.t.y {
		rope.t.y--
	}

	if rope.h.y > rope.t.y {
		rope.t.y++
	}
}

func (rope *Rope) moveLeft() {
	rope.h.x--

	if isCoordinateNear(rope.t, rope.h) {
		return
	}

	rope.t.x--

	if rope.h.y < rope.t.y {
		rope.t.y--
	}

	if rope.h.y > rope.t.y {
		rope.t.y++
	}
}

func (rope *Rope) moveUp() {
	rope.h.y = rope.h.y + 1

	if isCoordinateNear(rope.t, rope.h) {
		return
	}

	rope.t.y++

	if rope.h.x > rope.t.x {
		rope.t.x++
	}

	if rope.h.x < rope.t.x {
		rope.t.x--
	}
}

func (rope *Rope) moveDown() {
	rope.h.y--

	if isCoordinateNear(rope.t, rope.h) {
		return
	}

	rope.t.y--

	if rope.h.x > rope.t.x {
		rope.t.x++
	}

	if rope.h.x < rope.t.x {
		rope.t.x--
	}
}

func move(rope *Rope, howMany int, direction string) []GridCoordinate {
	tailCoordinates := []GridCoordinate{}

	for i := 0; i < howMany; i++ {
		if direction == "L" {
			rope.moveLeft()
		}

		if direction == "R" {
			rope.moveRight()
		}

		if direction == "U" {
			rope.moveUp()
		}

		if direction == "D" {
			rope.moveDown()
		}

		tailCoordinates = append(tailCoordinates, GridCoordinate{y: rope.t.y, x: rope.t.x})
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
	rope := Rope{h: GridCoordinate{0, 0}, t: GridCoordinate{0, 0}}
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

		allMoves = append(allMoves, move(&rope, numberOfSteps, direction)...)
	}

	uniqueMoves := make(map[string]struct{})

	for _, move := range allMoves {
		k := fmt.Sprintf("%d - %d", move.y, move.x)
		uniqueMoves[k] = struct{}{}
	}

	fmt.Println(len(uniqueMoves))

	readFile.Close()

}
