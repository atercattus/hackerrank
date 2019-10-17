package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Item struct {
	Cnt int
	Val int
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n     int
		s     string
		items [26]Item
	)
	fmt.Fscan(rdr, &n)
	fmt.Fscan(rdr, &s)

	totals := 0
	for _, ch := range s {
		ch := int(ch)
		items[ch-'a'].Cnt++
		cnt := items[ch-'a'].Cnt
		if cnt == 1 {
			totals++
			items[ch-'a'].Val = ch
		}
	}

	if totals < 2 {
		fmt.Fprintln(wrr, `0`)
		return
	}

	sort.Slice(items[:], func(i, j int) bool {
		return items[i].Cnt > items[j].Cnt
	})

	maxLen := 0
	for i := 0; i < totals; i++ {
		stay1 := items[i].Val
		for j := i + 1; j < totals; j++ {
			stay2 := items[j].Val
			curLen := 0
			prev := 0

			for _, ch := range s {
				ch := int(ch)
				if stay1 == ch || stay2 == ch {
					if ch == prev {
						curLen = 0
						break
					}
					curLen++
					prev = ch
					continue
				}
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

	fmt.Fprintln(wrr, maxLen)
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
