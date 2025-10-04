package linkedlist

import "fmt"

func TestDoubleLinkedList() {
	type Cat struct {
		name string
		age  uint8
	}

	dLinkedList := NewDoublyLinkedList[Cat]()
	dLinkedList.PushHead(Cat{"Max", 4})
	dLinkedList.PushHead(Cat{"Alex", 5})
	dLinkedList.PushTail(Cat{"Tom", 7})
	dLinkedList.Insert(2, Cat{"Tommy", 1})
	fmt.Printf("List size: %d\n", dLinkedList.Size())
	dLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
	data, _ := dLinkedList.Get(1)
	fmt.Printf("Get data with index: %d, %+v\n", 1, data)

	fmt.Println("--Remove data--")
	dLinkedList.Remove(3)
	dLinkedList.ForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})

	fmt.Println("--ReverseForEach--")
	dLinkedList.ReverseForEach(func(data Cat) {
		fmt.Printf("%+v\n", data)
	})
}
