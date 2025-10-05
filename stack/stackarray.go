package stack

import (
	"errors"
	"fmt"
)

type StackArray[T any] struct {
	fixedSize uint
	data      []T
}

func NewStackArray[T any](fixedSize uint) *StackArray[T] {
	stackArray := &StackArray[T]{
		fixedSize: fixedSize,
	}
	return stackArray
}

func (s *StackArray[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *StackArray[T]) Size() int {
	return len(s.data)
}

func (s *StackArray[T]) Push(value T) error {
	if uint(s.Size()) >= s.fixedSize {
		return errors.New("stack overflow")
	}
	s.data = append([]T{value}, s.data...)
	return nil
}

func (s *StackArray[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	pop := s.data[0]
	s.data = s.data[1:]
	return pop, nil
}

func (s *StackArray[T]) Peak() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return s.data[0], nil
}

func (s *StackArray[T]) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}
	fmt.Print("Stack: [")
	for _, it := range s.data {
		fmt.Printf("%+v, ", it)
	}
	fmt.Println("]")
}
