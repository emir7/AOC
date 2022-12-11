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
	c := 0

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		split := strings.Split(line, ",")
		elf1 := strings.Split(split[0], "-")
		elf2 := strings.Split(split[1], "-")

		v11, _ := strconv.Atoi(elf1[0])
		v12, _ := strconv.Atoi(elf1[1])

		v21, _ := strconv.Atoi(elf2[0])
		v22, _ := strconv.Atoi(elf2[1])

		if v11 < v21 && (v12 >= v21 && v12 <= v22) {
			c++
		} else if v11 > v21 && (v12 >= v21 && v12 <= v22) {
			c++
		} else if v21 < v11 && (v22 >= v11 && v22 <= v12) {
			c++
		} else if v21 > v11 && (v22 >= v11 && v22 <= v12) {
			c++
		} else if v11 <= v21 && v12 >= v22 {
			c++
		} else if v21 <= v11 && v22 >= v12 {
			c++
		}

	}

	fmt.Println("total: ", c)

	readFile.Close()
}
