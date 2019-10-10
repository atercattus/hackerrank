package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n int
	)
	fmt.Fscan(rdr, &n)

	fact := big.NewInt(1)
	for i := 2; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	fmt.Fprintln(wrr, fact)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
