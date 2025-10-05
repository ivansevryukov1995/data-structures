package binarysearchtree

import (
	"fmt"
	"math/rand"
)

func TestBinarySearchTree() {
	// rand.Seed(time.Now().Unix())
	workers := []*Worker{
		{"Max", 83}, {"Alex", 58}, {"Tom", 98}, {"Tommy", 62},
		{"Max", 70}, {"Alex", 34}, {"Tom", 22},
		{"Max", 60}, {"Alex", 99}, {"Tom", 91}, {"Tommy", 94},
		{"Tommy", 85},
	}
	tree := NewBinarySearchTree[*Worker]()
	for _, it := range workers {
		tree.Insert(it)
		fmt.Println("----------------------------")
		fmt.Printf("Added %+v\n", *it)
		tree.PrintTree()
	}
	fmt.Println("----- Find ------")
	key, err := tree.Find(62)
	fmt.Printf("Founded key value: %+v , err: %v\n", key, err)
	fmt.Println("----- Max key ------")
	key, err = tree.Maximum()
	fmt.Printf("Max key value: %+v , err: %v\n", key, err)
	fmt.Println("----- Min key ------")
	key, err = tree.Minimum()
	fmt.Printf("Min key value: %+v , err: %v\n", key, err)
	fmt.Println("----- Remove ------")
	removingKey := workers[rand.Intn(len(workers))]
	fmt.Printf("Remove node with key: %+v\n", *removingKey)
	tree.Remove(removingKey.GetKey())
	tree.PrintTree()
	// Варианты обхода дерева
	tree.SymmetricTraversal(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
	tree.TraversalAftertProcessing(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
	tree.TraversalBeforeProcessing(func(data *Worker) {
		fmt.Printf("%v ", data.GetKey())
	})
}
