package main

import (
	"fmt"
	"sort"
)

func solve() {
	var (
		n int
	)
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	// one fake point
	a = append(a, a[len(a)-1]+100)
	n++

	maxCnt := 0
	for min := 0; min < n; min++ {
		for max := min + 1; max < n; max++ {
			diff := a[max] - a[min]
			if diff > 1 {
				if cnt := max - min; cnt > maxCnt {
					maxCnt = cnt
				}
				break
			}
		}
	}

	fmt.Println(maxCnt)
}

func main() {
	solve()
}
