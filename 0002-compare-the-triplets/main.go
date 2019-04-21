package main

import (
	"fmt"
	"os"
)

func read(a *[3]int) {
	fmt.Fscanln(os.Stdin, &a[0], &a[1], &a[2])
}

func solve(a, b [3]int) (solution [2]int) {
	for i := 0; i < len(a); i++ {
		if diff := a[i] - b[i]; diff > 0 {
			solution[0]++
		} else if diff < 0 {
			solution[1]++
		}
	}
	return
}

func main() {
	var (
		alice, bob [3]int
	)

	read(&alice)
	read(&bob)
	solution := solve(alice, bob)
	fmt.Println(solution[0], solution[1])
}
