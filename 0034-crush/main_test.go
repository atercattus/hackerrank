package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
5 3
1 2 100
2 5 100
3 4 100
`, `
200
`)
}
