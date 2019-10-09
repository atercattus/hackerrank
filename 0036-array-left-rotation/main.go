package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, d int
	)
	fmt.Fscan(rdr, &n, &d)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(wrr, arr[(i+d)%n], ` `)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
