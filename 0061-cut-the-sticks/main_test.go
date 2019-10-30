package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
6
5 4 4 2 2 8
`, `
6
4
2
1
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
8
1 2 3 4 3 3 2 1
`, `
8
6
4
1
`)
}
