package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
2
4 3
-1 -3 4 2
4 2
0 -1 2 1
`, `
YES
NO
`)
}
