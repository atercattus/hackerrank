package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	// mine:
	hackerrank.Checker(solve, t, `
6
3 4 5 6 7 2
`, `
no
`)
}

func Test_solve1(t *testing.T) {
	// mine:
	hackerrank.Checker(solve, t, `
6
1 5 3 4 2 6
`, `
yes
swap 2 5
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
2  
4 2
`, `
yes
swap 1 2
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
3
3 1 2
`, `
no
`)
}

func Test_solve4(t *testing.T) {
	hackerrank.Checker(solve, t, `
6
1 5 4 3 2 6
`, `
yes
reverse 2 5
`)
}

func Test_solve5(t *testing.T) {
	hackerrank.Checker(solve, t, `
6
43 65 1 98 99 101
`, `
no
`)
}

func Test_solve6(t *testing.T) {
	hackerrank.Checker(solve, t, `
6
1 2 3 5 4 6
`, `
yes
swap 4 5
`)
}

func Test_solve7(t *testing.T) {
	// mine
	hackerrank.Checker(solve, t, `
6
1 2 3 6 5 4
`, `
yes
swap 4 6
`)
}
