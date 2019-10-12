package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func calc(mat [][]int) (sum int) {
	prev := mat[0]
	for r := 1; r < len(mat); r++ {
		cur := mat[r]
		for c, cv := range cur { // можно не просчитывать крайние стороны
			pv := prev[c]
			sum += abs(pv - cv)
		}
		prev = cur
	}
	return
}

func rotate(mat [][]int) [][]int {
	dst := make([][]int, len(mat[0]))
	for row := range mat[0] {
		dst[row] = make([]int, len(mat))
	}
	for col := range mat {
		for row := range mat[col] {
			dst[row][col] = mat[col][row]
		}
	}
	return dst
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		H, W int
	)
	fmt.Fscan(rdr, &H, &W)

	mat := make([][]int, H+2)
	for i := range mat {
		mat[i] = make([]int, W+2)
	}

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			fmt.Fscan(rdr, &mat[h+1][w+1])
		}
	}

	area := W * H * 2 // за верх и них
	area += calc(mat)
	mat = rotate(mat)
	area += calc(mat)

	fmt.Fprintln(wrr, area)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
