package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
10
beabeefeab
`, `
5
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
28
asdcbsdcagfsdbgdfanfghbsfdab
`, `
8
`)
}
