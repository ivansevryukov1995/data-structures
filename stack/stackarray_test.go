package stack

import (
	"fmt"
	"testing"
)

func TestStackArray(t *testing.T) {
	type Cat struct {
		name string
		age  uint8
	}
	var fSize uint = 3
	stack := NewStackArray[Cat](fSize)
	stack.Push(Cat{"Max", 4})
	stack.Push(Cat{"Alex", 5})
	stack.Push(Cat{"Tom", 7})
	stack.PrintStack()
	err := stack.Push(Cat{"Tommy", 1})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("----- Peak ------")
	peakValue, _ := stack.Peak()
	fmt.Printf("Stack head value: %+v\n", peakValue)
	stack.PrintStack()
	fmt.Println("----- Pop ------")
	for it := 0; it < int(fSize)+1; it++ {
		popValue, err := stack.Pop()
		fmt.Printf("Stack head value: %+v, err: %v\n", popValue, err)
		stack.PrintStack()
	}
}
