package main

import (
	"fmt"
	"os"
)

func frac(cur, tot int) {
	fmt.Printf("%.6f\n", float64(cur)/float64(tot))
}

func main() {
	var (
		cnt, num, pos, neg, zero int
	)

	fmt.Fscanln(os.Stdin, &cnt)
	for i := 1; i <= cnt; i++ {
		fmt.Fscan(os.Stdin, &num)

		if num > 0 {
			pos++
		} else if num < 0 {
			neg++
		} else {
			zero++
		}
	}

	frac(pos, cnt)
	frac(neg, cnt)
	frac(zero, cnt)
}
