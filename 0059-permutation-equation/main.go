package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n        int
		p, pm, y []int
	)
	fmt.Fscan(rdr, &n)

	p = make([]int, n+1)
	pm = make([]int, n+1)
	for i := 1; i <= n; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)
		p[i] = cur
		pm[cur] = i
	}

	y = make([]int, n+1)
	for i := 1; i <= n; i++ {
		idx := p[pm[i]]
		y[p[p[idx]]] = idx
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(wrr, y[i])
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
