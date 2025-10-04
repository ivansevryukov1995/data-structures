package dynamicarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	type Cat struct {
		name string
		age  uint8
	}

	dynamicArray := NewDynamicArray[Cat](1)
	dynamicArray.Add(Cat{"Max", 4})
	dynamicArray.Add(Cat{"Alex", 5})

	value1, _ := dynamicArray.Get(1)
	assert.Equal(t, "Alex", value1.name)
	assert.Equal(t, 2, dynamicArray.GetCapacity())

	dynamicArray.Add(Cat{"Tom", 7})

	assert.Equal(t, 4, dynamicArray.capacity)

	dynamicArray.Remove(0)
	dynamicArray.Put(1, Cat{"Tommy", 1})
}
