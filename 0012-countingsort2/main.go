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

	var cnts [100]int32

	for _, v := range arr {
		cnts[v]++
	}

	for i, cnt := range cnts {
		if cnt == 0 {
			continue
		}

		for r := int32(1); r <= cnt; r++ {
			fmt.Print(i, ` `)
		}
	}
	fmt.Println()
}
