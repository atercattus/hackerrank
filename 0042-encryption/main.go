package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		s string
	)
	fmt.Fscan(rdr, &s)

	sq := math.Sqrt(float64(len(s)))
	rows := int(math.Floor(sq))
	cols := int(math.Ceil(sq))
	if rows*cols < len(s) {
		rows++
	}

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			i := r*cols + c
			if i < len(s) {
				fmt.Fprint(wrr, string(s[i]))
			}
		}
		fmt.Fprint(wrr, ` `)
	}
	fmt.Fprintln(wrr)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
