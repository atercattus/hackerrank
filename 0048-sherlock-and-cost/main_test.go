package main

import (
	"github.com/atercattus/hackerrank"
	"io/ioutil"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
1
3
1 2 3
`, `
2
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
1
5
10 1 10 1 10
`, `
36
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
2
5
3 15 4 12 10
3
4 7 9
`, `
50
12
`)
}

func Test_solve4(t *testing.T) {
	in, _ := ioutil.ReadFile(`input14.txt`)
	out, _ := ioutil.ReadFile(`output14.txt`)
	hackerrank.Checker(solve, t, string(in), string(out))
}
