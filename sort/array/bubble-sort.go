package main

import (
	"fmt"
)

func main() {
	a := [5]int{5, 4, 3, 2, 1}
	b := []int{}
	b = append(b, a[:]...)
	fmt.Println(b)
	/* Bubble sort */
	for i := 0; i < len(b)-1; i++ {
		for j := 0; j < len(b)-i-1; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
	fmt.Println(b)

}

/*
Time Complexity: O(n^2)
Space Complexity: O(n)
*/
