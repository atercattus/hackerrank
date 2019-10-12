package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
haveaniceday
`, `
hae and via ecy
`)

	hackerrank.Checker(solve, t, `
feedthedog
`, `
fto ehg ee dd
`)

	hackerrank.Checker(solve, t, `
chillout
`, `
clu hlt io
`)
}
