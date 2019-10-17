package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Customer struct {
	Order    int
	PrepTime int
	Pos      int
}

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n         int
		customers []Customer
	)
	fmt.Fscan(rdr, &n)

	customers = make([]Customer, n)
	for i := 0; i < n; i++ {
		var order, prep int
		fmt.Fscan(rdr, &order, &prep)
		customers[i].Pos = i
		customers[i].Order = order
		customers[i].PrepTime = prep
	}

	sort.Slice(customers, func(i, j int) bool {
		if diff := (customers[i].PrepTime + customers[i].Order) - (customers[j].PrepTime + customers[j].Order); diff != 0 {
			return diff < 0
		}
		return customers[i].Order > customers[j].Order
	})

	var result []string
	for i := range customers {
		result = append(result, strconv.Itoa(customers[i].Pos+1))
	}
	fmt.Fprintln(wrr, strings.Join(result, ` `))
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
