package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n   int
		arr []int
	)
	fmt.Fscan(rdr, &n)

	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	minDiff := math.MaxInt32
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, abs(arr[i]-arr[i-1]))
	}

	fmt.Fprintln(wrr, minDiff)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
