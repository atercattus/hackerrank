package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(rdr *bufio.Reader, wrr *bufio.Writer) {
	var (
		b, n, m int
		keys    []int
		max     int = -1
	)
	fmt.Fscanln(rdr, &b, &n, &m)

	for i := 0; i < n; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)
		if cur < b {
			keys = append(keys, cur)
		}
	}

	for i := 0; i < m; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)

		for _, key := range keys {
			if sum := key + cur; sum <= b && sum > max {
				max = sum
			}
		}
	}

	fmt.Fprintln(wrr, max)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
