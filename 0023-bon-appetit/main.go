package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(rdr *bufio.Reader, wrr *bufio.Writer) {
	var (
		n, k        int
		expect, got int
	)
	fmt.Fscanln(rdr, &n, &k)

	for i := 0; i < n; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)
		if i != k {
			expect += cur
		}
	}
	expect /= 2

	fmt.Fscan(rdr, &got)

	if expect == got {
		fmt.Fprintln(wrr, `Bon Appetit`)
	} else {
		fmt.Fprintln(wrr, got-expect)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
