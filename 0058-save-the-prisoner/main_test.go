package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
2
5 2 1
5 2 2
`, `
2
3
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
1
4 6 2
`, `
3
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
2
7 19 2
3 7 3
`, `
6
3
`)
}
