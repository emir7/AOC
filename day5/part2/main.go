package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func appendToBeginning(dst []string, src []string) []string {
	final := []string{}

	for i := 0; i < len(src); i++ {
		final = append(final, src[i])
	}

	for i := 0; i < len(dst); i++ {
		final = append(final, dst[i])
	}

	return final
}

func part1() {

}

func main() {
	currentStack := [][]string{
		{"T", "Z", "B"},
		{"N", "D", "T", "H", "V"},
		{"D", "M", "F", "B"},
		{"L", "Q", "V", "W", "G", "J", "T"},
		{"M", "Q", "F", "V", "P", "G", "D", "W"},
		{"S", "F", "H", "G", "Q", "Z", "V"},
		{"W", "C", "T", "L", "R", "N", "S", "Z"},
		{"M", "R", "N", "J", "D", "W", "H", "Z"},
		{"S", "D", "F", "L", "Q", "M"},
	}

	/*currentStack := [][]string{
		{"N", "Z"},
		{"D", "C", "M"},
		{"P"},
	}*/
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

		splittedLine := strings.Split(line, " ")
		n, _ := strconv.Atoi(splittedLine[1])
		f, _ := strconv.Atoi(splittedLine[3])
		t, _ := strconv.Atoi(splittedLine[5])

		currentStack[t-1] = appendToBeginning(currentStack[t-1], currentStack[f-1][0:n])
		currentStack[f-1] = currentStack[f-1][n:]
	}

	str := ""

	for i := 0; i < len(currentStack); i++ {
		str = str + currentStack[i][0]
	}

	fmt.Println(str)
	readFile.Close()

}
