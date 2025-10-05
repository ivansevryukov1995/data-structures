package queue

import "fmt"

func TestQueueArray() {
	type Cat struct {
		name string
		age  uint8
	}
	var fSize uint = 3
	queue := NewQueueArray[Cat](fSize)
	queue.Enqueue(Cat{"Max", 4})
	queue.Enqueue(Cat{"Alex", 5})
	queue.Enqueue(Cat{"Tom", 7})
	queue.PrintQueue()
	err := queue.Enqueue(Cat{"Tommy", 1})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("----- Peak ------")
	peakValue, _ := queue.Peak()
	fmt.Printf("Queue head value: %+v\n", peakValue)
	queue.PrintQueue()
	fmt.Println("----- Dequeue ------")
	for it := 0; it < int(fSize)+1; it++ {
		popValue, err := queue.Dequeue()
		fmt.Printf("Queue head value: %+v, err: %v\n", popValue, err)
		queue.PrintQueue()
	}
}
