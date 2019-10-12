package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func out(a []int, wrr io.Writer) {
	var buf []string
	for _, v := range a {
		buf = append(buf, strconv.Itoa(v))
	}
	fmt.Fprintln(wrr, strings.Join(buf, ` `))
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t, n, k int
	)
	fmt.Fscan(rdr, &t)

	for testIdx := 0; testIdx < t; testIdx++ {
		fmt.Fscan(rdr, &n, &k)

		a := make([]int, n)
		if k == 0 {
			for i := range a {
				a[i] = i + 1
			}
		} else if n%(2*k) != 0 {
			fmt.Fprintln(wrr, -1)
			continue
		} else {
			for i := range a {
				b := i / k
				if b%2 == 0 {
					a[(b+1)*k+(i%k)] = i + 1
				} else {
					a[(b-1)*k+(i%k)] = i + 1
				}
			}
		}
		out(a, wrr)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
