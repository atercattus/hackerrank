package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
aba
10
`, `
7
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
a
1000000000000
`, `
1000000000000
`)
}
