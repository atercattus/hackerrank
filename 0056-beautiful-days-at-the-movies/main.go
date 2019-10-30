package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		i, j, k int
		cnt     int
	)
	fmt.Fscan(rdr, &i, &j, &k)

	for day := i; day <= j; day++ {
		rev := 0
		for d := day; d > 0; {
			m := d % 10
			d = d / 10
			rev = rev*10 + m
		}

		diff := day - rev
		if diff < 0 {
			diff = -diff
		}
		if diff%k == 0 {
			cnt++
		}
	}

	fmt.Fprintln(wrr, cnt)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
