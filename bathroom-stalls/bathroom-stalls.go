package main

import (
	"fmt"
	"math"
)

func main() {
	ncases := 0
	fmt.Scan(&ncases)
	for i := 0; i < ncases; i++ {
		start()
	}
}

func start() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = i + 1
		r[n-i-1] = i + 1
	}
	minIndex(l, r)
}

func minIndex(l, r []int) {
	m := min(l[0], r[0])
	index := 0
	for i := 0; i < len(l); i++ {
		if m > min(l[i], r[i]) {
			m = min(l[i], r[i])
			index = i
		}
	}
	fmt.Println(index, m)

}
func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
