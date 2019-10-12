package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
1 1
1
`, `
6
`)

	hackerrank.Checker(solve, t, `
3 3
1 3 4
2 2 3
1 2 4
`, `
60
`)

	hackerrank.Checker(solve, t, `
1 10
91 80 7 41 36 11 48 57 40 43
`, `
1276
`)
}
