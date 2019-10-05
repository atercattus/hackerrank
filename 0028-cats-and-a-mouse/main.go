package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(rdr *bufio.Reader, wrr *bufio.Writer) {
	var (
		q int
	)
	fmt.Fscanln(rdr, &q)

	abs := func(v int) int {
		if v < 0 {
			v = -v
		}
		return v
	}

	for i := 0; i < q; i++ {
		var x, y, z int
		fmt.Fscanln(rdr, &x, &y, &z)

		if diff := abs(x-z) - abs(y-z); diff == 0 {
			fmt.Fprintln(wrr, `Mouse C`)
		} else if diff < 0 {
			fmt.Fprintln(wrr, `Cat A`)
		} else {
			fmt.Fprintln(wrr, `Cat B`)
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
