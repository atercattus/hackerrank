package main

import (
	"bufio"
	"fmt"
	"os"
)

const MB = 1024 * 1024

func main() {
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	var (
		n, k int
		ar   []int
	)

	fmt.Fscanln(rdr, &n, &k)

	ar = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &ar[i])
	}

	count := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}

	fmt.Fprintln(wrr, count)
}
