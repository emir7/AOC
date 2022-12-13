package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	currentItems   []int
	operation      string
	divisibleBy    int
	caseT          int
	caseF          int
	inspectedItems int
}

func execOperation(n int, operation string) int {
	s := strings.Split(operation, " ")
	o := strings.Trim(s[1], " ")
	s0 := strings.Trim(s[0], " ")
	s2 := strings.Trim(s[2], " ")

	var v1 int
	var v2 int

	if s0 == "old" {
		v1 = n
	} else {
		v1, _ = strconv.Atoi(s0)
	}

	if s2 == "old" {
		v2 = n
	} else {
		v2, _ = strconv.Atoi(s2)

	}

	if o == "+" {
		return v1 + v2
	}

	return v1 * v2
}

func printMonkeys(monkeys []*Monkey) {
	for index, monkey := range monkeys {
		fmt.Println(monkey.currentItems, index)
	}
}

func (monkey *Monkey) removeItem(itemToBeRemoved int) {
	result := []int{}
	removedIndex := -1

	for i, item := range monkey.currentItems {
		if itemToBeRemoved == item {
			removedIndex = i
			break
		}
	}

	for i, v := range monkey.currentItems {
		if i == removedIndex {
			continue
		}

		result = append(result, v)
	}

	monkey.inspectedItems++
	monkey.currentItems = result
}

func (monkey *Monkey) addItem(item int) {
	monkey.currentItems = append(monkey.currentItems, item)
}

func play(monkeys []*Monkey) {
	numberOfRounds := 10000
	currentMonkeyIndex := 0

	bigMod := 1
	for _, m := range monkeys {
		bigMod *= m.divisibleBy
	}

	for numberOfRounds > 0 {
		currentMonkey := monkeys[currentMonkeyIndex]

		clonedItems := make([]int, len(currentMonkey.currentItems))
		copy(clonedItems, currentMonkey.currentItems)

		for _, item := range clonedItems {
			worryLevel := execOperation(item, currentMonkey.operation)

			if worryLevel < 0 {
				fmt.Println("whooooo", worryLevel)
			}

			worryLevel %= bigMod
			currentMonkey.removeItem(item)

			if worryLevel%currentMonkey.divisibleBy == 0 {
				monkeys[currentMonkey.caseT].addItem(worryLevel)
			} else {
				monkeys[currentMonkey.caseF].addItem(worryLevel)
			}
		}

		currentMonkeyIndex = (currentMonkeyIndex + 1) % len(monkeys)

		if currentMonkeyIndex == 0 {
			numberOfRounds--
		}
	}
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	monkeys := []*Monkey{}

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if err != nil {
			return
		}

		splittedLine := strings.Fields(line)
		var currentMonkey *Monkey

		if len(splittedLine) == 0 {
			continue
		}

		if splittedLine[0] == "Monkey" {
			monkeys = append(monkeys, &Monkey{})
			continue
		}

		currentMonkey = monkeys[len(monkeys)-1]

		if splittedLine[0] == "Starting" {
			splittedStartingItemsLine := strings.Split(line, ":")
			numbersStr := strings.Split(splittedStartingItemsLine[1], ",")

			items := []int{}

			for _, nStr := range numbersStr {
				v, _ := strconv.Atoi(strings.Trim(nStr, " "))
				items = append(items, v)
			}

			currentMonkey.currentItems = items
		}

		if splittedLine[0] == "Operation:" {
			colonLineSplit := strings.Split(line, "=")
			currentMonkey.operation = strings.Trim(colonLineSplit[1], " ")
		}

		if splittedLine[0] == "Test:" {
			lastIndex := len(splittedLine) - 1
			s := strings.Trim(splittedLine[lastIndex], " ")
			n, _ := strconv.Atoi(s)
			currentMonkey.divisibleBy = n
		}

		if splittedLine[0] == "If" && splittedLine[1] == "true:" {
			lastIndex := len(splittedLine) - 1
			s := strings.Trim(splittedLine[lastIndex], " ")
			toMonkey, _ := strconv.Atoi(s)
			currentMonkey.caseT = toMonkey
		}

		if splittedLine[0] == "If" && splittedLine[1] == "false:" {
			lastIndex := len(splittedLine) - 1
			s := strings.Trim(splittedLine[lastIndex], " ")
			toMonkey, _ := strconv.Atoi(s)
			currentMonkey.caseF = toMonkey
		}
	}

	play(monkeys)
	//fmt.Println(monkeys[0])
	finalArr := []int{}

	for _, monkey := range monkeys {
		finalArr = append(finalArr, monkey.inspectedItems)
	}

	fmt.Println(finalArr)
	sort.Ints(finalArr)
	//fmt.Println(finalArr)

	topMonkey := finalArr[len(finalArr)-1]
	penultimateMonkey := finalArr[len(finalArr)-2]

	solution := topMonkey * penultimateMonkey

	fmt.Println(solution)

	readFile.Close()

}
