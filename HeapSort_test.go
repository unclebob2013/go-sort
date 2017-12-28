package sort

import (
	"fmt"
	"testing"
)

func Test_heap_sort(t *testing.T) {
	array := []int{-1, 5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array[1:])
	HeapSort(array)
	fmt.Println("HeapSort:", array[1:])
}
