package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, m int
	)

	fmt.Fscan(rdr, &n, &m)

	arr := make([]int, n+1)

	for task := 0; task < m; task++ {
		var f, t, val int
		fmt.Fscan(rdr, &f, &t, &val)
		arr[f-1] += val
		arr[t] -= val
	}

	max, sum := 0, 0
	for _, v := range arr {
		sum += v
		if sum > max {
			max = sum
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
