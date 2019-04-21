package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MB = 1024 * 1024

type Kang struct {
	Pos  int32
	Jump int32
}

func main() {
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	var k1, k2 Kang

	fmt.Fscanln(rdr, &k1.Pos, &k1.Jump, &k2.Pos, &k2.Jump)

	ok := false

	if k1.Jump == k2.Jump {
		ok = k2.Pos == k1.Pos
	} else {
		n := float64(k2.Pos-k1.Pos) / float64(k1.Jump-k2.Jump)
		ok = n >= 0 && (math.Floor(n)-n) == 0
	}

	fmt.Fprintln(wrr, map[bool]string{false: `NO`, true: `YES`}[ok])
}
