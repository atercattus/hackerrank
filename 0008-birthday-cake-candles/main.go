package main

import (
	"fmt"
)

func main() {
	var (
		n, height         int
		maxHeight, maxCnt int
	)

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&height)
		if height > maxHeight {
			maxHeight = height
			maxCnt = 1
		} else if height == maxHeight {
			maxCnt++
		}
	}

	fmt.Println(maxCnt)
}
