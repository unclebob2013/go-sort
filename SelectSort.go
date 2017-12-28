package sort

import (
	"fmt"
)

func SelectSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		fmt.Println("i=", i)
		minIndex := i
		for j := i + 1; j < len(n); j++ {
			fmt.Println("j=", j)
			if n[j] < n[minIndex] {
				fmt.Println("minIndex", minIndex)
				minIndex = j
			}
			if minIndex != i {
				fmt.Println("swap: ", n[minIndex], " ", n[i])
				swapArray(n, minIndex, i)
			}
		}
	}
	return n
}
