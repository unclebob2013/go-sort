package sort

import (
	"fmt"
)

func InsertSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		fmt.Println("i=", i)
		for j := i; j > 0; j-- {
			fmt.Println("j=", j)
			if n[j-1] > n[j] {
				fmt.Println("swap: ", n[j-1], " ", n[j])
				swapArray(n, j-1, j)
			}
			fmt.Println(n)
		}
	}
	return n
}
