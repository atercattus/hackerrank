package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(n, p int) int {
	fromBegin := p / 2
	fromEnd := (n - p + ((n + 1) % 2)) / 2

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	return min(fromBegin, fromEnd)
}

func solve(rdr *bufio.Reader, wrr *bufio.Writer) {
	var (
		n, p int
	)
	fmt.Fscanln(rdr, &n)
	fmt.Fscanln(rdr, &p)
	fmt.Fprintln(wrr, calc(n, p))
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
