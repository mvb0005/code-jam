package main

import "fmt"

func main() {
	ncases := 0
	fmt.Scan(&ncases)
	for i := 0; i < ncases; i++ {
		fmt.Printf("Case #%d: %f\n", i+1, start())
	}
}

func start() float64 {
	d, n := 0, 0
	k, s := 0, 0
	fmt.Scan(&d, &n)
	max := 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&k, &s)
		if max < float64(d-k)/float64(s) {
			max = float64(d-k) / float64(s)
		}
	}
	return float64(d) / max

}
