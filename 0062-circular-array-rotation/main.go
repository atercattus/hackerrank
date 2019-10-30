package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, k, q int
		arr     []int
	)
	fmt.Fscan(rdr, &n, &k, &q)

	arr = make([]int, n)
	for i := range arr {
		fmt.Fscan(rdr, &arr[i])
	}

	for query := 0; query < q; query++ {
		var m int
		fmt.Fscan(rdr, &m)

		j := m - k
		for j < 0 {
			j += n
		}

		fmt.Fprintln(wrr, arr[j%n])
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
