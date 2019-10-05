package main

import (
	"bufio"
	"fmt"
	"os"
)

const MB = 1024 * 1024

func main() {
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)

	var (
		n     int
		types [6]int
	)

	fmt.Fscanln(rdr, &n)

	for i := 0; i < n; i++ {
		var id int
		fmt.Fscan(rdr, &id)
		types[id]++
	}

	maxId, maxCnt := 0, 0
	for id, cnt := range types {
		if cnt > maxCnt {
			maxCnt = cnt
			maxId = id
		}
	}

	fmt.Println(maxId)
}
