package queue

import "fmt"

func TestQueueLinkedList() {
	type Cat struct {
		name string
		age  uint8
	}
	queue := NewQueueLinkedList[Cat]()
	queue.Enqueue(Cat{"Max", 4})

	queue.Enqueue(Cat{"Alex", 5})
	queue.Enqueue(Cat{"Tom", 7})
	queue.Enqueue(Cat{"Tommy", 1})
	queue.PrintQueue()
	fmt.Println("----- Peak ------")
	peakValue, _ := queue.Peak()
	fmt.Printf("Queue head value: %+v\n", peakValue)
	queue.PrintQueue()
	fmt.Println("----- Dequeue ------")
	for !queue.IsEmpty() {
		popValue, err := queue.Dequeue()
		fmt.Printf("Queue head value: %+v, err: %v\n", popValue, err)
		queue.PrintQueue()
	}
}
