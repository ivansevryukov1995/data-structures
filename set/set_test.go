package set

import "fmt"

func main() {
	type Cat struct {
		name string
		age  uint8
	}
	setA := NewSet(1, 2, 3, 4, 5, 6, 3, 3, 3, 2, 1, 4, 7, 8, 5)
	fmt.Println("----- setA ------")
	setA.PrintSet()

	setB := NewSet(10, 22, 1, 4, 21, 4, 5, 21, 11, 10)
	fmt.Println("----- setB ------")
	setB.PrintSet()

	fmt.Println("----- Contains ------")
	value := 10 // поменяйте на любой существующий ключ
	fmt.Printf("SetA is contains'%v' == %v\n",
		value, setA.Contains(value))
	fmt.Printf("SetB is contains '%v'== %v\n",
		value, setB.Contains(value))

	fmt.Println("----- A - B ------")
	setA.Difference(setB).PrintSet()
	fmt.Println("----- B - A ------")
	setB.Difference(setA).PrintSet()
	fmt.Println("----- Symmetric Difference ------")
	setA.SymmetricDifference(setB).PrintSet()
	fmt.Println("----- Union ------")
	setA.Union(setB).PrintSet()
	fmt.Println("----- Intersect ------")
	setA.Intersect(setB).PrintSet()
}
