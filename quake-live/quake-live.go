package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, splitTeams())
	}

}

func splitTeams() int {
	var teamsize int
	fmt.Scan(&teamsize)
	var skillList = make([]int, teamsize)
	for i := 0; i < teamsize; i++ {
		fmt.Scan(&skillList[i])
	}
	sort.Ints(skillList)
	t1 := 0
	t2 := 0

	for i := 0; i < teamsize/2; i++ {
		if i%2 == 0 {
			t1 += skillList[i]
			t2 += skillList[i+teamsize/2]
		} else {
			t2 += skillList[i]
			t1 += skillList[i+teamsize/2]
		}
	}
	return int(math.Abs(float64(t1 - t2)))
}
