package main

import (
	"github.com/atercattus/hackerrank"
	"io/ioutil"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
3
1 3
2 3
3 3
`, `
1 2 3
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
5
8 1
4 2
5 6
3 1
4 3
`, `
4 2 5 1 3
`)
}

func Test_solve3(t *testing.T) {
	in, _ := ioutil.ReadFile(`input07.txt`)
	out, _ := ioutil.ReadFile(`output07.txt`)
	hackerrank.Checker(solve, t, string(in), string(out))
}
