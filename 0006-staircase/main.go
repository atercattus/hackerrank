package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		cnt int
	)

	fmt.Fscanln(os.Stdin, &cnt)
	for r := 1; r <= cnt; r++ {
		for c := 1; c <= cnt; c++ {
			if c <= (cnt - r) {
				fmt.Print(` `)
			} else {
				fmt.Print(`#`)
			}
		}
		fmt.Println()
	}
}
