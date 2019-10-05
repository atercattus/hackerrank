package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(rdr *bufio.Reader, wrr *bufio.Writer) {
	var (
		n int
	)
	fmt.Fscanln(rdr, &n)

	cnts := make([]int, 101)

	for i := 0; i < n; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)
		cnts[cur]++
	}

	pairs := 0
	for _, cnt := range cnts {
		pairs += (cnt >> 1)
	}

	fmt.Fprintln(wrr, pairs)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
