package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve0(t *testing.T) {
	hackerrank.Checker(solve, t, `
hackerhappy
hackerrank
9
`, `
Yes
`)
}

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
aba
aba
7
`, `
Yes
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
ashley
ash
2
`, `
No
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
aba
aba
2
`, `
Yes
`)
}

func Test_solve4(t *testing.T) {
	hackerrank.Checker(solve, t, `
zzzzz
zzzzzzz
4
`, `
Yes
`)
}
