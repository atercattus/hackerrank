package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
5
ab
bb
hefg
dhck
dkhc
`, `
ba
no answer
hegf
dhkc
hcdk
`)

	hackerrank.Checker(solve, t, `
6
lmno
dcba
dcbb
abdc
abcd
fedcbabcd
`, `
lmon
no answer
no answer
acbd
abdc
fedcbabdc
`)
}
