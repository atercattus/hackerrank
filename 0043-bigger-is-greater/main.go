package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		T int
	)
	fmt.Fscan(rdr, &T)

	for t := 0; t < T; t++ {
		var w []byte
		fmt.Fscan(rdr, &w)

		pos1 := len(w) - 2
		for pos1 >= 0 && w[pos1] >= w[pos1+1] {
			pos1--
		}
		if pos1 == -1 {
			fmt.Fprintln(wrr, `no answer`)
			continue
		}

		pos2 := len(w) - 1
		for pos2 >= 0 && w[pos2] <= w[pos1] {
			pos2--
		}
		w[pos1], w[pos2] = w[pos2], w[pos1]

		tail := w[pos1+1:]
		for i := 0; i < len(tail)/2; i++ {
			ni := len(tail) - i - 1
			tail[i], tail[ni] = tail[ni], tail[i]
		}

		fmt.Fprintln(wrr, string(w))
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
