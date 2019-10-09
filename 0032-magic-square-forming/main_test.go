package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	// mine:
	hackerrank.Checker(solve, t, `
8 3 4
1 5 9
6 7 2
`, `
0
`)

	// real:
	hackerrank.Checker(solve, t, `
5 3 4
1 5 8
6 4 2
`, `
7
`)

	hackerrank.Checker(solve, t, `
4 9 2
3 5 7
8 1 5
`, `
1
`)

	hackerrank.Checker(solve, t, `
4 8 2
4 5 7
6 1 6
`, `
4
`)
}
