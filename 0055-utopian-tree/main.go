package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t       int
		heights [61]int
	)
	fmt.Fscan(rdr, &t)

	heights[0] = 1
	for i := 1; i < len(heights); i++ {
		if (i-1)%2 == 0 {
			heights[i] = 2 * heights[i-1]
		} else {
			heights[i] = 1 + heights[i-1]
		}
	}

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(rdr, &n)
		fmt.Fprintln(wrr, heights[n])
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
