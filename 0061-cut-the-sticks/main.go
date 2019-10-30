package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

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

	cnt := len(arr)

	fmt.Fprintln(wrr, cnt)
	for p := 1; p < n; p++ {
		cnt--
		if arr[p] != arr[p-1] {
			fmt.Fprintln(wrr, cnt)
		}
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
