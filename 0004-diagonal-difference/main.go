package main

import (
	"fmt"
	"os"
)

func absDiff(a, b int64) (r int64) {
	if r = a - b; r < 0 {
		r = -r
	}
	return
}

func main() {
	var (
		cnt               int
		diag1, diag2, num int64
	)

	fmt.Fscanln(os.Stdin, &cnt)
	for r := 1; r <= cnt; r++ {
		for c := 1; c <= cnt; c++ {
			fmt.Fscan(os.Stdin, &num)

			if r == c {
				diag1 += num
			}
			if r == (cnt - c + 1) {
				diag2 += num
			}
		}
	}

	fmt.Println(absDiff(diag1, diag2))
}
