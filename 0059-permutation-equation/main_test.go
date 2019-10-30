package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
5
5 2 1 3 4
`, `
4
2
5
1
3
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
3
2 3 1
`, `
2
3
1
`)
}
