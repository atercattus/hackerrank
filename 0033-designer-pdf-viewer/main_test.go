package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5
abc
`, `
9
`)

	hackerrank.Checker(solve, t, `
1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 7
zaba
`, `
28
`)
}
