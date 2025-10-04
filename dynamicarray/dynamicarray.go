package dynamicarray

import (
	"errors"
	"fmt"
)

type DynamicArray[T any] struct {
	length   int
	capacity int
	arr      []T
}

func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	if capacity <= 0 {
		panic("Array capacity cannot be <= 0")
	}
	return &DynamicArray[T]{
		capacity: capacity,
		arr:      make([]T, capacity),
	}
}

func (da *DynamicArray[T]) GetLenght() int {
	return da.length
}

func (da *DynamicArray[T]) GetCapacity() int {
	return da.capacity
}

// Проверка на правильность обращения к элементу по индексу
func (da *DynamicArray[T]) checkRangeFromIndex(index int) error {
	if index >= da.length || index < 0 {
		return errors.New(fmt.Sprintf("Index %d out of range %d", index, da.length))
	}
	return nil
}

// Выделения увеличение размера памяти,выделяемого под массив
func (da *DynamicArray[T]) newCapacity() {
	da.capacity = da.capacity << 1
	newArr := make([]T, da.capacity)
	copy(newArr, da.arr)
	da.arr = newArr

}

func (da *DynamicArray[T]) IsEmpty() bool {
	return da.length == 0
}

// Добавление нового элемнта в массив
func (da *DynamicArray[T]) Add(element T) {
	if da.length == da.capacity {
		da.newCapacity()
	}

	da.arr[da.length] = element
	da.length++

}

// Удаление элемнта по индексу
func (da *DynamicArray[T]) Remove(index int) error {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return err
	}
	copy(da.arr[index:], da.arr[index+1:da.length])
	da.arr[da.length-1] = *new(T)
	da.length--

	return nil
}

// Получение занчения элемнта по индексу
func (da *DynamicArray[T]) Get(index int) (T, error) {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return *new(T), err
	}
	return da.arr[index], nil
}

// Обновление значения элемента массива
func (da *DynamicArray[T]) Put(index int, elenent T) error {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return err
	}
	da.arr[index] = elenent

	return nil
}
