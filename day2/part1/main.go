package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	winMap := make(map[string]string)
	crossMap := make(map[string]string)
	pointsMap := make(map[string]int)

	finalScore := [3]int{0, 3, 6}

	winMap["A"] = "Y"
	winMap["B"] = "Z"
	winMap["C"] = "X"

	crossMap["A"] = "X"
	crossMap["B"] = "Y"
	crossMap["C"] = "Z"

	pointsMap["X"] = 1
	pointsMap["Y"] = 2
	pointsMap["Z"] = 3

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

		splittedStrategy := strings.Fields(line)
		oponnetOption := splittedStrategy[0]
		myOption := splittedStrategy[1]

		sum += pointsMap[myOption]

		if myOption == crossMap[oponnetOption] {
			sum += finalScore[1]
		}

		if winMap[oponnetOption] == myOption {
			sum += finalScore[2]
		}
	}

	fmt.Println("total: ", sum)

	readFile.Close()
}
