package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
5 1
4 2 3 5 1
`, `
5 2 3 4 1
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
3 1
2 1 3
`, `
3 1 2
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
2 1
2 1
`, `
2 1
`)
}
