package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
4 3
1 2 3
`, `
4
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
10 4
2 5 3 6
`, `
5
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
222 24
3 25 34 38 26 42 16 10 15 50 39 44 36 29 22 43 20 27 9 30 47 13 40 33
`, `
5621927
`)
}
