package main

import (
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, k int
	)
	fmt.Fscanln(rdr, &n, &k)

	max := 0

	for i := 0; i < n; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)

		if diff := cur - k; diff > max {
			max = diff
		}
	}

	fmt.Fprintln(wrr, max)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
