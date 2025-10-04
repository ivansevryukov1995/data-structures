package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList(t *testing.T) {
	type Cat struct {
		name string
		age  uint8
	}

	sLinkedList := NewSingleLinkedList[Cat]()
	sLinkedList.PushHead(Cat{"Max", 4})
	sLinkedList.PushHead(Cat{"Alex", 5})
	sLinkedList.PushTail(Cat{"Tom", 7})
	sLinkedList.Insert(2, Cat{"Tommy", 1})

	value1, _ := sLinkedList.Get(3)
	assert.Equal(t, value1.name, "Tommy")
	assert.Equal(t, sLinkedList.Size(), 4)

	sLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	data, _ := sLinkedList.Get(1)
	fmt.Printf("Get data with index: %d, %+v\n", 1, data)

	fmt.Println("--Remove data--")
	sLinkedList.Remove(1)
	sLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
}
