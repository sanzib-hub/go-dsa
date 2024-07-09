package main

import (
	"fmt"
)

func main() {
	a := [5]int{5, 4, 3, 2, 1}
	b := []int{}
	b = append(b, a[:]...)
	fmt.Println(b)
	/* Selection sort */
	for i := 0; i < len(b)-1; i++ {
		index := i
		for j := i + 1; j < len(b); j++ {
			if b[j] < b[index] {
				index = j
			}
		}
		b[i], b[index] = b[index], b[i]
	}
	fmt.Println(b)

}

/*
Time Complexity: O(n^2)
Space Complexity: O(n)
*/
