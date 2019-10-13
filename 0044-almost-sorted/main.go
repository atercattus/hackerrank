package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(rdr io.Reader, wrr io.Writer) {
	var (
		n int
	)
	fmt.Fscan(rdr, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rdr, &arr[i])
	}
	arr = append(arr, 10000)

	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	check := func() bool {
		result := ``
		swapped := [2]int{-1, -1}
		locMin := -1
	loop:
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				for j := i + 1; j < n; j++ {
					if arr[j] > arr[j-1] {
						continue
					}

					if arr[j] < arr[j-1] && (j == n-1 || (j > i+1 && arr[j] < arr[j+1])) {
						swap(i, j)
						swapped = [2]int{i, j}
						result = fmt.Sprintf(`swap %d %d`, i+1, j+1)
						break loop
					} else if arr[j] < arr[j-1] && arr[j] < arr[j+1] {
						// Локальный минимум рядом с i, проверяемый отдельно
						locMin = j
					}
				}
			}
		}

		// Особая ситуация, когда своп между соседними элементами
		if swapped[1] == -1 && locMin > -1 {
			swap(locMin-1, locMin)
			swapped = [2]int{locMin - 1, locMin}
			result = fmt.Sprintf(`swap %d %d`, locMin, locMin+1)
		}

		// Проверка, что после swap последовательность стала корректной
		needReverseCheck := false
		if swapped[1] >= 0 {
			checkFrom := swapped[0] - 1
			if checkFrom < 0 {
				checkFrom = 0
			}
			for i := checkFrom; i < n-1; i++ {
				if arr[i] > arr[i+1] {
					needReverseCheck = true
					swap(swapped[0], swapped[1])
					break
				}
			}
		}

		// swap не чинит, пробую проверить на reverse
		if needReverseCheck {
			f, t := swapped[0], swapped[1]
			mid := arr[f : t+1]
			l := len(mid)

			for i := 0; i < l/2; i++ {
				ni := l - i - 1
				swap(f+i, f+ni)
			}

			checkFrom := f - 1
			if checkFrom < 0 {
				checkFrom = 0
			}
			for i := checkFrom; i < n-1; i++ {
				if arr[i] > arr[i+1] {
					return false
				}
			}
			result = fmt.Sprintf(`reverse %d %d`, f+1, t+1)
		}

		if result != `` {
			fmt.Fprintln(wrr, "yes\n"+result+"\n")
			return true
		}
		return false
	}

	if !check() {
		fmt.Fprintln(wrr, `no`)
	}
}

func main() {
	const MB = 1024 * 1024
	rdr := bufio.NewReaderSize(os.Stdin, 4*MB)
	wrr := bufio.NewWriterSize(os.Stdout, 4*MB)
	defer wrr.Flush()

	solve(rdr, wrr)
}
