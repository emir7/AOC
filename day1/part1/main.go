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
	fileScanner.Split(bufio.ScanLines)

	max := 0
	currentNum := 0

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			fmt.Println(currentNum)
			if currentNum > max {
				max = currentNum
			}

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
	fmt.Println(max)
}
