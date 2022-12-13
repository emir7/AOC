package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	currentCycle := 0
	currentX := 1
	interestingCycle := 20
	sum := 0

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

				if currentCycle == interestingCycle {
					sum += currentCycle * currentX
					interestingCycle += 40
				}

				currentCycle++
				v, _ := strconv.Atoi(splittedLine[1])

				if currentCycle == interestingCycle {
					sum += currentCycle * currentX
					interestingCycle += 40
				}

				currentX += v
			}
		} else {
			if splittedLine[0] == "noop" {
				currentCycle++

				if currentCycle == interestingCycle {
					sum += currentCycle * currentX
					interestingCycle += 40
				}
			}
		}
	}

	fmt.Println(sum)

	readFile.Close()

}
