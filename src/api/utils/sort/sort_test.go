package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	if len(elements) != 10 {
		t.Error("slice should contains 10 elements")
	}

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Error("Bubble sort took more than 500ms")
		return
	}

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0], "first ele should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last ele should be 9")
}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "first ele should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last ele should be 9")
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
