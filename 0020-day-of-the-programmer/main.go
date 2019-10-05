package main

import (
	"bufio"
	"fmt"
	"os"
)

func dayOfProgrammer(year int) string {
	var leap bool
	var day int

	if year == 1918 {
		day = 26
	} else {
		if year < 1918 {
			leap = (year%4 == 0)
		} else {
			leap = (year%400 == 0) || (year%4 == 0 && year%100 != 0)
		}

		day = 13
		if leap {
			day--
		}
	}

	return fmt.Sprintf(`%02d.09.%d`, day, year)
}

const MB = 1024 * 1024

func main() {
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	var (
		year int
	)

	fmt.Fscanln(rdr, &year)

	fmt.Fprintln(wrr, dayOfProgrammer(year))
}
