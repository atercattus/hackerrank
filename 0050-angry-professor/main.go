package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t, n, k int
	)
	fmt.Fscan(rdr, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(rdr, &n, &k)

		onTime := 0
		for j := 0; j < n; j++ {
			var cur int
			fmt.Fscan(rdr, &cur)
			if cur <= 0 {
				onTime++
			}
		}

		if onTime < k {
			fmt.Fprintln(wrr, `YES`)
		} else {
			fmt.Fprintln(wrr, `NO`)
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
