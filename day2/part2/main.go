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
	loseMap := make(map[string]string)
	crossMap := make(map[string]string)
	pointsMap := make(map[string]int)
	resultMap := make(map[string]string)

	finalScore := [3]int{0, 3, 6}

	resultMap["X"] = "lose"
	resultMap["Y"] = "draw"
	resultMap["Z"] = "win"

	winMap["A"] = "Y"
	winMap["B"] = "Z"
	winMap["C"] = "X"

	loseMap["A"] = "Z"
	loseMap["B"] = "X"
	loseMap["C"] = "Y"

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
		result := splittedStrategy[1]

		if resultMap[result] == "lose" {
			playedSymbol := loseMap[oponnetOption]
			sum += finalScore[0]
			sum += pointsMap[playedSymbol]
		}

		if resultMap[result] == "draw" {
			playedSymbol := crossMap[oponnetOption]
			sum += finalScore[1]
			sum += pointsMap[playedSymbol]
		}

		if resultMap[result] == "win" {
			playedSymbol := winMap[oponnetOption]
			sum += finalScore[2]
			sum += pointsMap[playedSymbol]
		}

	}

	fmt.Println("total: ", sum)

	readFile.Close()
}
