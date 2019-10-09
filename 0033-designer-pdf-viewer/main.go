package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		chars [26]int
		str   string
	)

	for i := range chars {
		fmt.Fscan(rdr, &chars[i])
	}
	fmt.Fscan(rdr, &str)

	maxH := 0
	for i := range str {
		ch := str[i] - 'a'
		if h := chars[ch]; h > maxH {
			maxH = h
		}
	}

	fmt.Fprintln(wrr, maxH*len(str))
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
