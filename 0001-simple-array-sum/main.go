package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		cnt, sum, num int
	)

	fmt.Fscanln(os.Stdin, &cnt)
	if cnt > 0 {
		for i := 1; i <= cnt; i++ {
			fmt.Fscan(os.Stdin, &num)
			sum += num
		}
	}
	fmt.Println(sum)
}
