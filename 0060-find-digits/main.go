package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t, n int
	)
	fmt.Fscan(rdr, &t)

	for test := 1; test <= t; test++ {
		fmt.Fscan(rdr, &n)

		divs := 0
		for num := n; num > 0; num = num / 10 {
			digit := num % 10
			if digit == 0 {
				continue
			}
			if n%digit == 0 {
				divs++
			}
		}

		fmt.Fprintln(wrr, divs)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
