package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func moveCtr(str *[]rune, realX int) []rune {
	result := []rune("........................................")

	if realX < 0 {
		howMany := 3 + realX

		for i := 0; i < howMany; i++ {
			result[i] = []rune("#")[0]
		}

		return result
	}

	for i := realX; i < getMin(len(*str), realX+3); i++ {
		result[i] = []rune("#")[0]
	}

	return result
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	currentCycle := 0
	currentX := 1
	interestingCycle := 40
	sum := 0

	resetState := "###....................................."
	result := []rune("........................................")
	currentState := []rune(resetState)
	index := 0

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		splittedLine := strings.Fields(line)

		if len(splittedLine) == 2 {
			if splittedLine[0] == "addx" {
				currentCycle++

				//currentState = moveCtr(&currentState, currentX-1)
				result[index] = currentState[index]
				index++

				if currentCycle == interestingCycle {
					sum += currentCycle * currentX
					interestingCycle += 40
					index = 0
					fmt.Println(string(result))
				}

				currentCycle++
				//currentState = moveCtr(&currentState, currentX-1)
				result[index] = currentState[index]

				index++

				v, _ := strconv.Atoi(splittedLine[1])

				if currentCycle == interestingCycle {
					sum += currentCycle * currentX
					interestingCycle += 40
					index = 0
					fmt.Println(string(result))
				}

				currentX += v
				currentState = moveCtr(&currentState, currentX-1)
			}
		} else {
			if splittedLine[0] == "noop" {
				currentCycle++

				result[index] = currentState[index]
				index++

				if currentCycle == interestingCycle {
					sum += currentCycle * currentX
					interestingCycle += 40
					index = 0
					fmt.Println(string(result))
				}
			}
		}
	}

	readFile.Close()

}
