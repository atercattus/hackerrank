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
		n          int
		scores     []int
		min, max   int
		incr, decr int
	)

	fmt.Fscanln(rdr, &n)
	scores = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &scores[i])
	}

	min, max = scores[0], scores[0]

	for i := 1; i < len(scores); i++ {
		if score := scores[i]; max < score {
			incr++
			max = score
		} else if min > score {
			decr++
			min = score
		}
	}

	fmt.Fprintln(wrr, incr, decr)
}
