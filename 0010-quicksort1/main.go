package main

import (
	"fmt"
)

func divide(arr []int) {
	pos := 0
	pivot := arr[pos]
	minCnt := 0

	// how many elements are before pivot
	for _, v := range arr {
		if v < pivot {
			minCnt++
		}
	}

	if minCnt == 0 {
		return
	}

	// move pivot to centre
	arr[pos], arr[minCnt] = arr[minCnt], arr[pos]

	// move smaller elements to the left of array
	for i, v := range arr {
		if v < pivot {
			arr[i], arr[pos] = arr[pos], arr[i]
			pos++
		}
	}
}

func main() {
	var (
		n   int
		arr []int
	)

	fmt.Scanln(&n)
	arr = make([]int, n)

	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	divide(arr)

	for _, v := range arr {
		fmt.Print(v, ` `)
	}
	fmt.Println()
}
