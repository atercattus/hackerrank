package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t, n int
		B    []int
	)
	fmt.Fscan(rdr, &t)

	for ti := 0; ti < t; ti++ {
		fmt.Fscan(rdr, &n)
		B = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(rdr, &B[i])
			B[i]--
		}

		pair1, pair2 := 0, 0
		for i := 1; i < n; i++ {
			pair1New := max(pair1, pair2+B[i-1])
			pair2New := max(pair2, pair1+B[i])
			pair1, pair2 = pair1New, pair2New
		}
		fmt.Fprintln(wrr, max(pair1, pair2))
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
