package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isDuplicate(info []int) bool {
	c := 0

	for _, value := range info {
		if value > 0 {
			c++
		}
	}

	return c == 3
}

func findDuplicates(lines []string) (rune, bool) {
	duplicatesMap := make(map[rune][]int)

	for index, line := range lines {
		runeArr := []rune(line)
		for _, runeElement := range runeArr {
			_, ok := duplicatesMap[runeElement]

			if !ok {
				duplicatesMap[runeElement] = []int{0, 0, 0}
				duplicatesMap[runeElement][index]++
			} else {
				duplicatesMap[runeElement][index]++
			}

			if isDuplicate(duplicatesMap[runeElement]) {
				return runeElement, true
			}

		}
	}

	return rune(0), false
}

func getValue(char rune) int {
	isUpperCase := unicode.IsUpper(char)

	if isUpperCase {
		return int(char) - 65 + 27
	}

	return int(char) - 97 + 1
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	sum := 0
	lines := []string{}

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		lines = append(lines, line)

		if len(lines) == 3 {
			value, ok := findDuplicates(lines)

			if ok {
				sum += getValue(value)
			}

			lines = lines[:0]
		}

	}

	fmt.Println("total: ", sum)

	readFile.Close()
}
