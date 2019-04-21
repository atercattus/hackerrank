package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MB = 1024 * 1024

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	var (
		n, m int
		a, b []int

		min, max = math.MaxInt32, 0
		cnt      int
	)

	fmt.Fscanln(rdr, &n, &m)
	a = make([]int, n)
	b = make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &a[i])
		max = Max(max, a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(rdr, &b[i])
		min = Min(min, b[i])
	}

loop:
	for val := max; val <= min; val++ {
		for _, v := range a {
			if val%v != 0 {
				continue loop
			}
		}

		for _, v := range b {
			if v%val != 0 {
				continue loop
			}
		}

		cnt++
	}

	fmt.Fprintln(wrr, cnt)
}
