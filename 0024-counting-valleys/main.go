package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(rdr *bufio.Reader, wrr *bufio.Writer) {
	var (
		n, level, valleys int
	)
	fmt.Fscanln(rdr, &n)

	for i := 0; i < n; i++ {
		var cur byte
		cur, _ = rdr.ReadByte()
		if cur == 'U' {
			level++
		} else {
			level--
		}
		if level == 0 && cur == 'U' {
			valleys++
		}
	}

	fmt.Fprintln(wrr, valleys)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
