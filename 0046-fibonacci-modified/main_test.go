package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
0 1 5
`, `
5
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
0 1 6
`, `
27
`)
}
