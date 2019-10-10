package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var mem = map[string][]string{}

func solveReq(passwords [][]byte, attempt []byte) (result []string) {
	if resp, ok := mem[string(attempt)]; ok {
		return resp
	}

	for _, pass := range passwords {
		if !bytes.HasPrefix(attempt, pass) {
			continue
		}
		if tail := attempt[len(pass):]; len(tail) == 0 {
			result = append(result, string(pass))
			break
		} else if resp := solveReq(passwords, tail); resp != nil {
			result = append(result, string(pass))
			result = append(result, resp...)
			break
		}
	}

	mem[string(attempt)] = result

	return result
}

func solveOne(passwords [][]byte, attempt []byte, wrr io.Writer) {
	mem = map[string][]string{}

	result := solveReq(passwords, attempt)
	if result == nil {
		fmt.Fprintln(wrr, `WRONG PASSWORD`)
		return
	}
	fmt.Fprintln(wrr, strings.Join(result, ` `))
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		t int
	)
	fmt.Fscan(rdr, &t)

	for i := 0; i < t; i++ {
		var n int
		var attempt []byte
		fmt.Fscan(rdr, &n)
		passwords := make([][]byte, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(rdr, &passwords[j])
		}
		fmt.Fscan(rdr, &attempt)

		solveOne(passwords, attempt, wrr)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
