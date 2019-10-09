package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Rank struct {
	Score int
	Pos   int
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, m int
	)
	fmt.Fscan(rdr, &n)

	ranks := make([]Rank, n+1)
	prevScore := math.MaxInt32
	pos := 0
	for i := 0; i < n; i++ {
		var score int
		fmt.Fscan(rdr, &score)

		if score < prevScore {
			prevScore = score
			pos++
		}
		ranks[i] = Rank{Score: score, Pos: pos}
	}
	ranks[n] = Rank{Score: 0, Pos: pos + 1} // Чтобы не париться от границе

	fmt.Fscan(rdr, &m)

	pos = n
	for i := 0; i < m; i++ {
		var score int
		fmt.Fscan(rdr, &score)

		for (pos >= 0) && (ranks[pos].Score < score) {
			pos--
		}
		if pos < 0 {
			pos = 0
		}
		if ranks[pos].Score > score {
			pos++
		}
		fmt.Fprintln(wrr, ranks[pos].Pos)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
