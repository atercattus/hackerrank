package main

import (
	"fmt"
)

func main() {
	var (
		h, m, s int
		tz      string
		isAM    bool
	)

	fmt.Scanf(`%d:%d:%d%s`, &h, &m, &s, &tz)

	isAM = tz == `AM`
	if h == 12 {
		h = 0
	}

	if !isAM {
		h += 12
	}

	fmt.Printf("%02d:%02d:%02d\n", h, m, s)
}
