package main

import (
	"github.com/atercattus/hackerrank"
	"io/ioutil"
	"testing"
)

func Test_solve1(t *testing.T) {
	hackerrank.Checker(solve, t, `
2
4
1 2 3 4
6
2 -1 2 3 4 -5
`, `
10 10
10 11
`)
}

func Test_solve2(t *testing.T) {
	hackerrank.Checker(solve, t, `
1
5
-2 -3 -1 -4 -6
`, `
-1 -1
`)
}

func Test_solve3(t *testing.T) {
	hackerrank.Checker(solve, t, `
6
1
1
6
-1 -2 -3 -4 -5 -6
2
1 -2
3
1 2 3
1
-10
6
1 -1 -1 -1 -1 5
`, `
1 1
-1 -1
1 1
6 6
-10 -10
5 6
`)
}

func Test_solve4(t *testing.T) {
	in, _ := ioutil.ReadFile(`input01.txt`)
	out, _ := ioutil.ReadFile(`output01.txt`)
	hackerrank.Checker(solve, t, string(in), string(out))
}
