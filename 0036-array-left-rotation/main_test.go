package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
5 2
1 2 3 4 5
`, `
3 4 5 1 2
`)

	hackerrank.Checker(solve, t, `
5 4
1 2 3 4 5
`, `
5 1 2 3 4
`)
}
