package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Node struct {
	next  *Node
	prev  *Node
	value *[]interface{}
}

type Stack struct {
	h *Node
	t *Node
}

func (stack *Stack) add(el *[]interface{}) {
	newNode := &Node{next: nil, value: el}

	if stack.h == nil && stack.t == nil {
		stack.h = newNode
		stack.t = stack.h
		return
	}

	stack.t.next = newNode
	newNode.prev = stack.t
	stack.t = newNode
}

func (stack *Stack) remove() {
	if stack.t == stack.h {
		stack.h = nil
		stack.t = nil
		return
	}

	stack.t = stack.t.prev
	stack.t.next = nil
}

func (stack *Stack) print() {
	current := stack.h

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func parseLine(line []rune) *[]interface{} {
	stack := &Stack{h: nil, t: nil}

	for i := 0; i < len(line); i++ {
		if i == len(line)-1 {
			return stack.h.value
		}
		if line[i] == '[' {
			newList := []interface{}{}

			if stack.h != nil {
				*stack.t.value = append(*stack.t.value, &newList)
				stack.add(&newList)
			} else {
				stack.add(&newList)
			}

		} else if line[i] == ']' {
			stack.remove()
		} else if line[i] == ',' {
			continue
		} else {
			n, _ := strconv.Atoi(string(line[i]))
			*stack.t.value = append(*stack.t.value, n)
		}
	}

	return stack.h.value
}

func TypeSwitch(param interface{}) string {
	fmt.Println(reflect.TypeOf(param))
	switch param.(type) {

	case int:
		return "Its an int"
	case *[]interface{}:
		return "Its a slice"
	default:
		return "unlucky"
	}
}

func main() {
	/*readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	*/
	/*for fileScanner.Scan() {
		var line string = fileScanner.Text()
		fmt.Println(parseLine([]rune(line), [][]int{}, []int{}, 0))
	}*/

	parsedArr := parseLine([]rune("[1,[2,[3,4,5],[6,7,8]]]"))
	/*
		v, ok := secondPart.([]*interface{})
		fmt.Println(v, ok)
	*/
	secondPart := (*parsedArr)[0]
	fmt.Println(TypeSwitch(secondPart))
	//readFile.Close()

}
