package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		s, t string
		k    int
	)
	fmt.Fscan(rdr, &s, &t, &k)

	prefixLen := 0
	for i := 0; i < len(s) && i < len(t); i++ {
		if s[i] == t[i] {
			prefixLen++
		} else {
			break
		}
	}

	tailsLen := (len(s) - prefixLen) + (len(t) - prefixLen)
	if diff := k - tailsLen; diff >= 0 && diff%2 == 0 {
		fmt.Fprintln(wrr, `Yes`)
	} else if len(s)+len(t) <= k {
		fmt.Fprintln(wrr, `Yes`)
	} else {
		fmt.Fprintln(wrr, `No`)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
