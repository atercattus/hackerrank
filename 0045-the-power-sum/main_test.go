package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
10
2
`, `
1
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
100
2
`, `
3
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
100
3
`, `
1
`)
}
