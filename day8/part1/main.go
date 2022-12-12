package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMax(arr []string) int {
	max := -1

	for _, value := range arr {
		v, _ := strconv.Atoi(value)
		if v > max {
			max = v
		}
	}

	return max
}

func isTreeVisible(grid *[][]string, y int, x int) bool {
	if y == 0 || y == len(*grid)-1 {
		return true
	}

	if x == 0 || x == len((*grid)[y])-1 {
		return true
	}

	treeHeight, _ := strconv.Atoi((*grid)[y][x])

	// left
	mLeft := findMax((*grid)[y][0:x])

	if mLeft < treeHeight {
		//fmt.Println("left", mLeft)
		return true
	}

	// right
	mRight := findMax((*grid)[y][(x + 1):])

	if mRight < treeHeight {
		//fmt.Println("right")
		return true
	}

	// top
	topArr := []string{}

	for i := 0; i < y; i++ {
		topArr = append(topArr, (*grid)[i][x])
	}

	mTop := findMax(topArr)

	if mTop < treeHeight {
		return true
	}

	// bottom
	botArr := []string{}
	for i := y + 1; i < len(*grid); i++ {
		botArr = append(botArr, (*grid)[i][x])
	}

	mBot := findMax(botArr)

	if mBot < treeHeight {
		return true
	}

	return false
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	grid := [][]string{}

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		grid = append(grid, []string{})
		index := len(grid) - 1
		grid[index] = append(grid[index], strings.Split(line, "")...)
	}

	c := 0
	duplicate := make([][]string, len(grid))
	for i := range grid {
		duplicate[i] = make([]string, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len((grid)[i]); j++ {
			if isTreeVisible(&grid, i, j) {
				duplicate[i][j] = "X"
				fmt.Println(i, j, grid[i][j])
				c++
			}
		}
	}

	for i := 0; i < len(duplicate); i++ {
		fmt.Println(duplicate[i])
	}

	fmt.Println(c)

	//fmt.Println(isTreeVisible(&grid, 1, 1))

	readFile.Close()

}
