package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(current *[]int) int {
	sum := 0

	for _, element := range *current {
		sum += element
	}

	return sum
}

func updateMax(current *[]int, number int) *[]int {
	minIndex := 0
	minValue := math.MaxInt

	for index, element := range *current {
		if element < minValue {
			minIndex = index
			minValue = element
		}
	}

	if number > minValue {
		(*current)[minIndex] = number
	}

	return current
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	topThree := []int{0, 0, 0}
	currentNum := 0

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			updateMax(&topThree, currentNum)

			currentNum = 0
			continue
		}

		calories, err := strconv.Atoi(line)

		if err != nil {
			return
		}

		currentNum += calories

	}

	readFile.Close()
	fmt.Println(sum(&topThree))
}
