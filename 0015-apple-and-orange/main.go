package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MB = 1024 * 1024

func solve(rdr io.Reader, s, t, tree, fruits int) (cnt int) {
	for i := 0; i < fruits; i++ {
		var pos int
		fmt.Fscan(rdr, &pos)

		if pos += tree; s <= pos && pos <= t {
			cnt++
		}
	}
	return
}

func main() {
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	var s, t, a, b, m, n int
	fmt.Fscanln(rdr, &s, &t) // house
	fmt.Fscanln(rdr, &a, &b) // trees
	fmt.Fscanln(rdr, &m, &n) // a number of apples and oranges

	apples := solve(rdr, s, t, a, m)
	oranges := solve(rdr, s, t, b, n)

	fmt.Fprintln(wrr, apples)
	fmt.Fprintln(wrr, oranges)
}
