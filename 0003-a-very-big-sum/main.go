package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		cnt      int
		sum, num int64
	)

	fmt.Fscanln(os.Stdin, &cnt)
	for i := 1; i <= cnt; i++ {
		fmt.Fscan(os.Stdin, &num)
		sum += num
	}

	fmt.Println(sum)
}
