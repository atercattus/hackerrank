package main

import (
	"github.com/atercattus/hackerrank"
	"io/ioutil"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
3
2 1
3 0
3 2
`, `
2 1
1 2 3
-1
`)

	hackerrank.Checker(solve, t, `
10
2 1
10 5
7 5
2 1
2 0
2 0
1 0
10 5
10 0
6 0
`, `
2 1
6 7 8 9 10 1 2 3 4 5
-1
2 1
1 2
1 2
1
6 7 8 9 10 1 2 3 4 5
1 2 3 4 5 6 7 8 9 10
1 2 3 4 5 6
`)

	in, _ := ioutil.ReadFile(`input12.txt`)
	out, _ := ioutil.ReadFile(`output12.txt`)
	hackerrank.Checker(solve, t, string(in), string(out))
}
