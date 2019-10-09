package main

import (
	"github.com/atercattus/hackerrank"
	"testing"
)

func Test_solve(t *testing.T) {
	hackerrank.Checker(solve, t, `
4
aba
baba
aba
xzxb
3
aba
xzxb
ab
`, `
2
1
0
`)

	hackerrank.Checker(solve, t, `
3
def
de
fgh
3
de
lmn
fgh
`, `
1
0
1
`)

	hackerrank.Checker(solve, t, `
13
abcde
sdaklfj
asdjf
na
basdn
sdaklfj
asdjf
na
asdjf
na
basdn
sdaklfj
asdjf
5
abcde
sdaklfj
asdjf
na
basdn
`, `
1
3
4
3
2
`)
}
