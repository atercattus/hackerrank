package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	Item struct {
		Val int32
		Str string
	}

	SortItem struct {
		Cnt  int32
		Pos  int32
		Strs []string
	}
)

func load() (arr []Item) {
	var n, nHalf int

	rdr := bufio.NewReaderSize(os.Stdin, 4*1024*1024)

	fmt.Fscanln(rdr, &n)
	nHalf = n / 2
	arr = make([]Item, n)

	for i, _ := range arr {
		item := &arr[i]
		fmt.Fscanln(rdr, &item.Val, &item.Str)

		if i < nHalf {
			item.Str = `-`
		}
	}

	return
}

func main() {
	arr := load()

	var cnts [100]SortItem
	for _, v := range arr {
		cnts[v.Val].Cnt++
		cnts[v.Val].Strs = append(cnts[v.Val].Strs, v.Str)
	}

	wrr := bufio.NewWriterSize(os.Stdout, 4*1024*1024)

	for _, cnt := range cnts {
		if cnt.Cnt == 0 {
			continue
		}

		for _, str := range cnt.Strs {
			fmt.Fprint(wrr, str, ` `)
		}
	}
	fmt.Fprintln(wrr)
	wrr.Flush()
}
