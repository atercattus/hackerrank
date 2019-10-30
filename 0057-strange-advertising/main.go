package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, likes int
	)
	fmt.Fscan(rdr, &n)

	people := 5
	for ; n > 0; n-- {
		newPeople := people / 2
		likes += newPeople
		people = 3 * newPeople
	}

	fmt.Fprintln(wrr, likes)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
