package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMarkerIndex(input string) int {
	for i, v1 := range input {
		hm := make(map[rune]struct{})
		hm[v1] = struct{}{}
		isMarker := true

		for j := i + 1; j < getMin(len(input), i+14); j++ {
			_, ok := hm[rune(input[j])]

			if ok {
				isMarker = false
				break
			}

			hm[rune(input[j])] = struct{}{}
		}

		if isMarker {
			return i + 14
		}
	}

	return -1
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		fmt.Println(findMarkerIndex(line))
	}

	readFile.Close()

}
