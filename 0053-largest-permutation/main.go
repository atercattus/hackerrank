package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, k int
		arr  []int
	)
	fmt.Fscan(rdr, &n, &k)

	arr = make([]int, n+1)
	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var cur int
		fmt.Fscan(rdr, &cur)
		arr[i] = cur
		pos[cur] = i
	}

	max := n
	for k > 0 && max > 0 {
		wantPos := n - max + 1
		if gotPos := pos[max]; gotPos != wantPos {
			swap := arr[wantPos]
			arr[gotPos], arr[wantPos] = arr[wantPos], arr[gotPos]
			pos[swap], pos[max] = pos[max], pos[swap]
			k--
		}
		max--
	}

	var result []string
	for i := 1; i <= n; i++ {
		result = append(result, strconv.Itoa(arr[i]))
	}
	fmt.Fprintln(wrr, strings.Join(result, ` `))
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
