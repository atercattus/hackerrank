package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		s    string
		slen uint64
		n    uint64
	)
	fmt.Fscan(rdr, &s, &n)
	slen = uint64(len(s))

	cnt := uint64(0)
	for _, char := range s {
		if char == 'a' {
			cnt++
		}
	}

	cnt *= n / slen

	if head := n % slen; head != 0 {
		for i := uint64(0); i < head; i++ {
			if s[i] == 'a' {
				cnt++
			}
		}
	}

	fmt.Fprintln(wrr, cnt)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
