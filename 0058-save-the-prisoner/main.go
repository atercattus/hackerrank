package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t, n, m, s int
	)
	fmt.Fscan(rdr, &t)

	for test := 0; test < t; test++ {
		fmt.Fscan(rdr, &n, &m, &s)
		s = (s - 1) + m
		p := s % n
		if p == 0 {
			p = n
		}
		fmt.Fprintln(wrr, p)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
