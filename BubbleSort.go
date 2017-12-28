package sort

import (
	"fmt"
)

func checkArray(array []int) {
	if array == nil {
		panic("Nil array.")
	} else if len(array) <= 0 {
		panic("Empty array.")
	}
}

func swapArray(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func SimpleBubbleSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		fmt.Println("i=", i)
		for j := 1; j < len(n)-i; j++ {
			fmt.Println("j=", j)
			if n[j-1] > n[j] {
				swapArray(n, j-1, j)
			}
		}
	}
	return n
}

func FlagSwapBubbleSort(n []int) []int {
	var flag bool
	for i := 0; i < len(n); i++ {
		fmt.Println("i=", i)
		flag = false
		for j := 1; j < len(n)-i; j++ {
			fmt.Println("j=", j)
			if n[j-1] > n[j] {
				swapArray(n, j-1, j)
				flag = true
			}
		}
		if !flag {
			fmt.Println("no change break")
			break
		}
	}
	return n
}

func FlagSwapPositionBubbleSort(n []int) []int {
	var flag bool
	var swap int
	last_swap_position := len(n)

	for i := 0; i < len(n); i++ {
		fmt.Println("i=", i)
		// After the last swapping at x, [x, n] is sorted.
		// So we just need to sort [0, x] in the next scanning.
		flag = false
		swap = last_swap_position
		for j := 1; j < swap; j++ {
			fmt.Println("j=", j)
			if n[j-1] > n[j] {
				swapArray(n, j-1, j)
				flag = true
				last_swap_position = j
			}
		}
		if !flag {
			fmt.Println("no change break")
			break
		}
	}
	return n

}
