/*
	Google Codejam 2017 Worlds Problem Dice Straights
	I was able to solve up until the part where you need to verify that
	the straight uses differnt dices. The solution involves Bipartite
	Matching.
*/

package main

import (
	"fmt"
	"sort"
)

var (
	numberMap = make(map[int][]int)
	numDie    = 0
)

func main() {
	input()
	keys := getKeyList()
	k := getStraight(keys)
	fmt.Println(keys, k)
}

func input() {
	fmt.Scanln(&numDie)
	var temp int
	for i := 0; i < numDie; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&temp)
			numberMap[temp] = append(numberMap[temp], i)
		}
	}
}

func getKeyList() []int {
	var keys = make([]int, len(numberMap))
	counter := 0
	for i := range numberMap {
		keys[counter] = i
		counter++
	}
	sort.Ints(keys)
	return keys
}

func getStraight(keys []int) int {
	max := 0
	cur := -1
	for i := 0; i < len(keys)-1; i++ {
		fmt.Println(numberMap[keys[i]])
		//This requires Bipartite Matching https://en.wikipedia.org/wiki/Ford%E2%80%93Fulkerson_algorithm
		if keys[i]+1 == keys[i+1] {

		} else {
			if i-cur > max {
				max = i - cur
			}
			cur = i
		}
	}
	return max
}
