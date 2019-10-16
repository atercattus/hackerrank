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
		arr  []int
	)
	fmt.Fscan(rdr, &t)

	const minVal = -10000

	for i := 0; i < t; i++ {
		fmt.Fscan(rdr, &n)

		maxSubSeq := 0
		maxSubArr := minVal
		maxItem := minVal

		arr = arr[:0]

		accum := 0
		for j := 0; j < n; j++ {
			var cur int
			fmt.Fscan(rdr, &cur)
			arr = append(arr, cur)
			accum = max(cur, cur+accum)
			if cur > 0 {
				maxSubSeq += cur
			}
			maxSubArr = max(maxSubArr, accum)
			maxItem = max(maxItem, cur)
		}
		if maxSubSeq == 0 {
			maxSubSeq = maxItem
		}

		fmt.Fprintln(wrr, maxSubArr, maxSubSeq)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
