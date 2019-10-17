package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
3
3 -7 0
`, `
3
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
10
-59 -36 -13 1 -53 -92 -2 -96 -54 75
`, `
1
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
5
1 -3 71 68 17
`, `
3
`)
}
