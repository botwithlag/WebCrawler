package main

import (
	"fmt"
)

type Stack[T any] []T

type AnyNumb[T any] map[int]T

func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

func (stack *Stack[T]) Pop() T {
	if len(*stack) == 0 {
		var zero T
		return zero
	}
	value := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return value
}

func main() {
	var s Stack[string]
	s.Push("hello")
	s.Push("There")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
