package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

const magic = 15

func isMagic(sq [9]int) bool {
	return (magic == (sq[0] + sq[1] + sq[2])) &&
		(magic == (sq[3] + sq[4] + sq[5])) &&
		(magic == (sq[6] + sq[7] + sq[8])) &&
		(magic == (sq[0] + sq[3] + sq[6])) &&
		(magic == (sq[1] + sq[4] + sq[7])) &&
		(magic == (sq[2] + sq[5] + sq[8])) &&
		(magic == (sq[0] + sq[4] + sq[8])) &&
		(magic == (sq[2] + sq[4] + sq[6]))
}

var allMagicSquares [][9]int

func permutation(sq [9]int, from int, to int) {
	if from == to {
		if isMagic(sq) {
			allMagicSquares = append(allMagicSquares, sq)
		}
		return
	}
	for i := from; i <= to; i++ {
		sq[from], sq[i] = sq[i], sq[from]
		permutation(sq, from+1, to)
		sq[from], sq[i] = sq[i], sq[from]
	}
}

func diff(a, b [9]int) int {
	d := 0
	for i := range a {
		delta := a[i] - b[i]
		if delta < 0 {
			delta = -delta
		}
		d += delta
	}
	return d
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		sq [9]int
	)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			fmt.Fscan(rdr, &sq[r*3+c])
		}
	}

	mat := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	permutation(mat, 0, 8)

	min := math.MaxInt32
	for _, mat := range allMagicSquares {
		d := diff(sq, mat)
		if d < min {
			min = d
		}
	}

	fmt.Fprintln(wrr, min)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
