package main

import (
	"github.com/atercattus/hackerrank"
	"io/ioutil"
	"testing"
)

func Test_solve(t *testing.T) {
	// mine:
	hackerrank.Checker(solve, t, `
1
2
aaaaaaaaaa aaaaaaaaa
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
`, `
aaaaaaaaa aaaaaaaaa aaaaaaaaa aaaaaaaaa
`)

	hackerrank.Checker(solve, t, `
1
2
aaaaaaaaaa aaaaaaaaa
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	`, `
aaaaaaaaa aaaaaaaaa aaaaaaaaa aaaaaaaaa aaaaaaaaa aaaaaaaaa aaaaaaaaa aaaaaaaaa
	`)

	// real:
	hackerrank.Checker(solve, t, `
3
6
because can do must we what
wedowhatwemustbecausewecan
2
hello planet
helloworld
3
ab abcd cd
abcd
`, `
we do what we must because we can
WRONG PASSWORD
ab cd
`)

	hackerrank.Checker(solve, t, `
3
4
ozkxyhkcst xvglh hpdnb zfzahm
zfzahm
4
gurwgrb maqz holpkhqx aowypvopu
gurwgrb
10
a aa aaa aaaa aaaaa aaaaaa aaaaaaa aaaaaaaa aaaaaaaaa aaaaaaaaaa
aaaaaaaaaab
`, `
zfzahm
gurwgrb
WRONG PASSWORD
`)

	hackerrank.Checker(solve, t, `
3
4
ozkxyhkcst xvglh hpdnb zfzahm
zfzahm
4
gurwgrb maqz holpkhqx aowypvopu
gurwgrb
10
a aa aaa aaaa aaaaa aaaaaa aaaaaaa aaaaaaaa aaaaaaaaa aaaaaaaaaa
aaaaaaaaaab
`, `
zfzahm
gurwgrb
WRONG PASSWORD
	`)

	hackerrank.Checker(solve, t, `
4
6
zsnpgbqh bktrpgdwbu qhuhzxfh mxrgmga omgtgnqomb dffttrwlfh
nktrsgtwbuzsbptzahomgtgnaoma
6
xkof medbc mhyewjzsdg vkuzym tbeqv xtbyhn
xtbyhnmedbcmhyewjzsdgxtbyhn
6
alutwfal kkhbqlrxnm qyyx lwdgpchwic rdtgzvw sduh
sduhkkhbqlrxnmsduhsduhqyyx
10
a aa aaa aaaa aaaaa aaaaaa aaaaaaa aaaaaaaa aaaaaaaaa aaaaaaaaaa
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaab
`, `
WRONG PASSWORD
xtbyhn medbc mhyewjzsdg xtbyhn
sduh kkhbqlrxnm sduh sduh qyyx
WRONG PASSWORD
	`)

	in, _ := ioutil.ReadFile(`input08.txt`)
	out, _ := ioutil.ReadFile(`output08.txt`)
	hackerrank.Checker(solve, t, string(in), string(out))

	// Тесты ниже не проходят локальную проверку,
	//   потому что мой код выбирает другие подходящие варианты.
	// Но проверяльщик на сайте принимает решение.

	//in, _ = ioutil.ReadFile(`input26.txt`)
	//out, _ = ioutil.ReadFile(`output26.txt`)
	//hackerrank.Checker(solve, t, string(in), string(out))

	//in, _ = ioutil.ReadFile(`input27.txt`)
	//out, _ = ioutil.ReadFile(`output27.txt`)
	//hackerrank.Checker(solve, t, string(in), string(out))

	//in, _ = ioutil.ReadFile(`input32.txt`)
	//out, _ = ioutil.ReadFile(`output32.txt`)
	//hackerrank.Checker(solve, t, string(in), string(out))
}
