package sort

import (
	"fmt"
)

func quickSort(n []int, left int, right int) {
	low := left
	high := right
	pivot := n[(left+right)/2]
	fmt.Println("pivot=", pivot)
	for low <= high {
		fmt.Println("low:", low)
		fmt.Println("high:", high)
		// In the left part of the array, looking for the element which is greater than the pivot.
		for n[low] < pivot {
			low++
			fmt.Println("low change:", low)
		}
		// In the right part of the array, looking for the element which is lesser than the pivot.
		for n[high] > pivot {
			high--
			fmt.Println("high change:", high)
		}

		if low <= high {
			// Swap the two elements we found to meke sure that,
			// all values before 'low' element are less or equal than the pivot,
			// and all values after 'high' element are greater or equal to the pivot.
			fmt.Println("swap: ", n[low], " ", n[high])
			swapArray(n, low, high)
			low++
			high--
		}
	}

	// Handle the left part of the array which is unsorted.
	if left < high {
		fmt.Println("left part", "left=", left, " high=", high)
		fmt.Println(n)
		quickSort(n, left, high)
	}

	// Handle the right part of the array which is unsorted.
	if low < right {
		fmt.Println("right part", "low=", low, " right", right)
		fmt.Println(n)
		quickSort(n, low, right)
	}
}

func QuickSort(n []int) {
	quickSort(n, 0, len(n)-1)
}
