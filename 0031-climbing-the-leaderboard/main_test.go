package main

import (
	"github.com/atercattus/hackerrank"
	"io/ioutil"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
7
100 100 50 40 40 20 10
4
5 25 50 120
`, `
6
4
2
1
`)

	hackerrank.Checker(solve, t, `
6
100 90 90 80 75 60
5
50 65 77 90 102
`, `
6
5
4
2
1
`)

	in, _ := ioutil.ReadFile(`input07.txt`)
	out, _ := ioutil.ReadFile(`output07.txt`)
	hackerrank.Checker(solve, t, string(in), string(out))
}
