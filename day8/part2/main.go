package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countVisibleTreesArr(arr []string, height int) int {
	c := 0

	for _, value := range arr {
		v, _ := strconv.Atoi(value)
		if v < height {
			c++
		} else if v >= height {
			return c + 1
		}
	}

	return c
}

func countVisibleTrees(grid *[][]string, y int, x int) int {
	treeHeight, _ := strconv.Atoi((*grid)[y][x])

	// right
	cRight := countVisibleTreesArr((*grid)[y][x+1:], treeHeight)

	// left
	leftArr := []string{}

	for i := x - 1; i >= 0; i-- {
		leftArr = append(leftArr, (*grid)[y][i])
	}
	cLeft := countVisibleTreesArr(leftArr, treeHeight)

	// top
	topArr := []string{}

	for i := y - 1; i >= 0; i-- {
		topArr = append(topArr, (*grid)[i][x])
	}

	cTop := countVisibleTreesArr(topArr, treeHeight)
	// bottom
	botArr := []string{}
	for i := y + 1; i < len(*grid); i++ {
		botArr = append(botArr, (*grid)[i][x])
	}

	cBot := countVisibleTreesArr(botArr, treeHeight)

	return cLeft * cRight * cTop * cBot
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

	max := -1
	for i := range grid {
		for j := range grid[i] {
			p := countVisibleTrees(&grid, i, j)

			if p > max {
				max = p
			}
		}
	}

	fmt.Println(max)
	readFile.Close()

}
