package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n, m  int
		coins []int
	)
	fmt.Fscan(rdr, &n, &m)
	coins = make([]int, m)
	for i := range coins {
		fmt.Fscan(rdr, &coins[i])
	}

	var (
		mem      = map[[2]int]int{}
		solveReq func(needSum, coinIdx int) int
	)
	solveReq = func(needSum, coinIdx int) (found int) {
		if needSum == 0 {
			return 1
		} else if needSum < 0 {
			return 0
		} else if coinIdx >= len(coins) {
			return 0
		}

		if cnt, ok := mem[[2]int{needSum, coinIdx}]; ok {
			return cnt
		}

		for idx := coinIdx; idx < len(coins); idx++ {
			coin := coins[idx]
			for cnt := 1; cnt <= needSum/coin; cnt++ {
				tailCnt := solveReq(needSum-coin*cnt, idx+1)
				found += tailCnt
			}
		}

		mem[[2]int{needSum, coinIdx}] = found

		return found
	}

	cnt := solveReq(n, 0)

	fmt.Fprintln(wrr, cnt)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
