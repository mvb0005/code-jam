/*
	Code Jam 2010 Qualification Snapper Chain
	This solution was derived by descovering that the bitset of
	power state was being iterated by 1, meaning that if the smallest
	n digits of k were all equal to 1, the power at n would be ON.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var ncases int
	fmt.Scan(&ncases)
	for casenum := 1; casenum <= ncases; casenum++ {
		var n, k int
		fmt.Scan(&n, &k)
		var o string
		if snap(n, k) {
			o = "ON"
		} else {
			o = "OFF"
		}
		fmt.Printf("Case #%d: %s\n", casenum, o)
	}

}

func snap(n, k int) bool {
	bitmask := int(math.Exp2(float64(n))) - 1
	return k&bitmask == bitmask
}
