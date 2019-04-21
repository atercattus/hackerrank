package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		numbers [5]int
		sum     int64
		sums    [5]int64
	)

	for i, _ := range numbers {
		fmt.Scan(&numbers[i])
		sum += int64(numbers[i])
	}

	for i, v := range numbers {
		sums[i] = sum - int64(v)
	}

	sort.Slice(sums[:], func(i, j int) bool {
		return sums[i] < sums[j]
	})

	fmt.Println(sums[0], sums[4])
}
