package linkedlist

import (
	"fmt"
	"testing"
)

func TestCycleLinkedList(t *testing.T) {
	type Cat struct {
		name string
		age  uint8
	}
	cLinkedList := NewCyclicLinkedList[Cat]()
	cLinkedList.Add(Cat{"Max", 4})
	cLinkedList.Add(Cat{"Alex", 5})
	cLinkedList.Add(Cat{"Tom", 7})
	cLinkedList.Add(Cat{"Tommy", 1})
	fmt.Printf("List size: %d\n", cLinkedList.Size())
	cLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	fmt.Println("--Remove data--")
	cLinkedList.Remove()
	cLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	fmt.Println("--Rotate--")
	data, _ := cLinkedList.Value()
	fmt.Printf("Head data before rotate: %+v\n", data)
	cLinkedList.Rotate(-1)
	data, _ = cLinkedList.Value()
	fmt.Printf("Head data after rotate: %+v\n", data)
	fmt.Println("--ReverseForEach--")
	cLinkedList.ReverseForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
}
