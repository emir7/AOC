package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func createMapFromFirstPart(part1 string) map[rune]struct{} {
	runeMap := make(map[rune]struct{})
	runeArr := []rune(part1)

	for _, value := range runeArr {
		runeMap[value] = struct{}{}
	}

	return runeMap
}

func findSecondPartDuplicates(part2 string, currentMap map[rune]struct{}) (rune, bool) {
	runeArr := []rune(part2)

	for _, element := range runeArr {
		_, ok := currentMap[element]

		if ok {
			return element, true
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

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		part1 := line[0 : len(line)/2]
		part2 := line[len(line)/2:]
		charMap := createMapFromFirstPart(part1)
		duplicate, ok := findSecondPartDuplicates(part2, charMap)

		if ok {
			sum += getValue(duplicate)
		}
	}

	fmt.Println("total: ", sum)

	readFile.Close()
}
