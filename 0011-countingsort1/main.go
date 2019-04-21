package main

import (
	"fmt"
)

func load() (arr []int) {
	var n int

	fmt.Scanln(&n)
	arr = make([]int, n)

	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	return
}

func main() {
	arr := load()

	var cnts [100]int
	for _, v := range arr {
		cnts[v]++
	}

	for _, v := range cnts {
		fmt.Print(v, ` `)
	}
	fmt.Println()
}
