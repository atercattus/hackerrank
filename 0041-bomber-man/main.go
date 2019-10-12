package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		r, c, n int
	)
	fmt.Fscan(rdr, &r, &c, &n)

	const bombDelay = 3

	grid := make([][]byte, r)
	for y := range grid {
		grid[y] = make([]byte, c)
		var line string
		fmt.Fscan(rdr, &line)
		for x := range grid[y] {
			ts := byte(0)
			if line[x] == 'O' {
				ts = bombDelay
			}
			grid[y][x] = ts
		}
	}

	// simulation
	// 0, 1, 2, (1, 0, 1, 2)*

	clear := func(x, y int) {
		if x < 0 || x >= len(grid[0]) {
			return
		}
		if y < 0 || y >= len(grid) {
			return
		}
		grid[y][x] = 0
	}

	if n > 3 {
		n = 3 + (n-3)%4
	}

	triggered := [][2]int{}
	for step := 0; step < n; step++ {
		triggered = triggered[:0]
		isPlantStep := step%2 == 1

		for y := range grid {
			for x := range grid[y] {
				cell := grid[y][x]
				if cell == 0 {
					if isPlantStep {
						grid[y][x] = bombDelay
					}
				} else {
					cell--
					grid[y][x] = cell
					if cell == 0 {
						triggered = append(triggered, [2]int{x, y})
					}
				}
			}
		}

		for _, xy := range triggered {
			x, y := xy[0], xy[1]
			clear(x, y)
			clear(x-1, y)
			clear(x+1, y)
			clear(x, y-1)
			clear(x, y+1)
		}
	}

	// print
	chars := map[bool]byte{true: 'O', false: '.'}
	for y := range grid {
		for x := range grid[y] {
			cell := grid[y][x]
			fmt.Fprintf(wrr, `%c`, chars[cell > 0])
		}
		fmt.Fprintln(wrr)
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
