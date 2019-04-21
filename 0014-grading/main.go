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

	var n int
	fmt.Fscanln(rdr, &n)

	for i := 0; i < n; i++ {
		var grade int
		fmt.Fscanln(rdr, &grade)

		if next := 5 * ((grade + 4) / 5); next-grade < 3 && next > 38 {
			grade = next
		}

		fmt.Fprintln(wrr, grade)
	}
	fmt.Fprintln(wrr)
}
