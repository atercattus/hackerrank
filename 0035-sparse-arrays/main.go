package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type STRING [2]uint64

func packString(str []byte) (s STRING) {
	for i := 0; i < 10 && i < len(str); i++ {
		ch := str[i] - 'a' + 1
		s[0] = (s[0] << 5) | uint64(ch)
	}

	for i := 10; i < 20 && i < len(str); i++ {
		ch := str[i] - 'a' + 1
		s[1] = (s[1] << 5) | uint64(ch)
	}

	return
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, q    int
		strings []STRING
	)
	fmt.Fscan(rdr, &n)

	strings = make([]STRING, n)
	for i := range strings {
		var str string
		fmt.Fscan(rdr, &str)
		strings[i] = packString([]byte(str))
	}

	fmt.Fscan(rdr, &q)

	for i := 0; i < q; i++ {
		var str string
		fmt.Fscan(rdr, &str)
		s := packString([]byte(str))
		cnt := 0
		for j := range strings {
			if s == strings[j] {
				cnt++
			}
		}
		fmt.Fprintln(wrr, cnt)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
