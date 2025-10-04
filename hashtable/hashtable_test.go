package hashmap

import (
	"fmt"
	"log"
	"testing"
)

func TestHashTable(t *testing.T) {
	type Cat struct {
		name string
		age  uint8
	}

	hashTable, _ := NewHashTableWithCapacity[string, Cat](3)
	hashTable.Put("Home", Cat{"Alex", 4})
	hashTable.Put("1", Cat{"Tom", 6})
	hashTable.Put("2", Cat{"Tommy", 6})
	hashTable.Put("2", Cat{"Max", 1}) // перезапись значения по ключу
	hashTable.Put("3", Cat{"Alex", 4})
	hashTable.Put("4", Cat{"Tom", 6})
	hashTable.Put("5", Cat{"Tommy", 6})
	hashTable.Put("Work", Cat{"Max", 1})
	hashTable.ForEach(func(key string, value Cat) {
		fmt.Printf("%v:%+v\n", key, value)
	})
	fmt.Println("-----Get Value------")
	val, err := hashTable.Get("1") // поменяйте на 0
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%+v\n", val)
	}
	fmt.Println("----- Remove ------")
	hashTable.Remove("Work")
	hashTable.Remove("5")
	hashTable.ForEach(func(key string, value Cat) {
		fmt.Printf("%v:%+v\n", key, value)
	})
	fmt.Println("----- Keys ------")
	fmt.Print("[")
	for _, it := range hashTable.Keys() {
		fmt.Printf("%+v  ", it)
	}
	fmt.Println("]")
	fmt.Println("----- Values ------")
	fmt.Print("[")
	for _, it := range hashTable.Values() {
		fmt.Printf("%+v  ", it)
	}
	fmt.Println("]")
	fmt.Println("----- Contains ------")
	key := "Home" // поменяйте на любой существующий ключ
	fmt.Printf("Key '%v' is contains == %v\n",
		key, hashTable.Contains(key))
	fmt.Println("----- Remove all ------")
	hashTable.Clear()
	hashTable.ForEach(func(key string, value Cat) {
		fmt.Printf("%v:%+v\n", key, value)
	})
}
