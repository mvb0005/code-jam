/*
	Google Code Jam 2018 Practice Round Number Guessing
*/

package main

import (
	"fmt"
)

func main() {
	ncases := 0
	fmt.Scan(&ncases)
	for i := 0; i < ncases; i++ {
		guess()
	}
}

func guess() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	n := 0
	fmt.Scan(&n)
	binarySearch(a, b)
}

func binarySearch(a, b int) {
	guess := int((float64(a+b) / 2.0))
	fmt.Println(guess)
	response := ""
	fmt.Scan(&response)
	if response == "CORRECT" {
		return
	} else if response == "TOO_BIG" {
		binarySearch(a, guess)
	} else if response == "TOO_SMALL" {
		if b-guess == 1 {
			fmt.Println(b)
			fmt.Scan(&response)
			return
		}
		binarySearch(guess, b)
	} else if response == "WRONG_ANSWER" {
		panic("we failed")
	}
}
