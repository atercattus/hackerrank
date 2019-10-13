package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solveReq(pow, from, to, maxSum int) int {
	if from > to {
		return 0
	} else if maxSum < 0 {
		return 0
	} else if maxSum == 0 {
		return 1
	}

	sum := int(math.Pow(float64(from), float64(pow)))
	if sum == maxSum {
		return 1
	}

	cntWo := solveReq(pow, from+1, to, maxSum)
	cntWith := solveReq(pow, from+1, to, maxSum-sum)

	return cntWith + cntWo
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		X, N int
	)
	fmt.Fscan(rdr, &X, &N)

	if X == 1 {
		fmt.Fprintln(wrr, 1)
		return
	}

	mid := int(math.Sqrt(float64(X)))

	cnt := solveReq(N, 1, mid, X)

	fmt.Fprintln(wrr, cnt)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
