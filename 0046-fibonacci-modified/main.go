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
		t1_, t2_, n int
	)
	fmt.Fscan(rdr, &t1_, &t2_, &n)

	t1 := big.NewInt(int64(t1_))
	t2 := big.NewInt(int64(t2_))
	ti := big.NewInt(0)

	for i := 3; i <= n; i++ {
		ti.Mul(t2, t2).Add(ti, t1)
		t1.Set(t2)
		t2.Set(ti)
	}

	fmt.Fprintln(wrr, ti.String())
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
