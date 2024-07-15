package main

import (
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
}
