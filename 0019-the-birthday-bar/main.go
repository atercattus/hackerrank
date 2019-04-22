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
		squares    []int
		day, month int
		count      int
	)

	fmt.Fscanln(rdr, &n)

	squares = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &squares[i])
	}
	fmt.Fscanln(rdr) // read eof

	fmt.Fscanln(rdr, &day, &month)

	for i := 0; i <= len(squares)-month; i++ {
		sum := 0
		for j := 0; j < month; j++ {
			if sum += squares[i+j]; sum > day {
				break
			}
		}
		if sum == day {
			count++
		}
	}

	fmt.Fprintln(wrr, count)
}
