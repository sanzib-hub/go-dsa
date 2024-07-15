//Order-agnostic binary search
package main

import "fmt"


func main(){
	a:=[6]int{2,9,17,29,47,91}
	b:=[6]int{91,47,29,17,9,2}

	fmt.Println("For a",binary(a[:], 29))
	fmt.Println("For v",binary(b[:], 29))

}

func binary(a []int, target int) int{
	start := 0
	end := len(a)-1
	var isAscending bool = a[start] < a[end]

	for start <= end{
		mid:= start + (end-start)/2
		if a[mid] == target{
			return mid
		}
		if a[mid] > target{
			if isAscending{
				end = mid -1
			}else{
				start = mid +1
			}
		} else if a[mid] < target{
			if isAscending{
				start = mid +1
			}else{
				end = mid -1
			}
		}
	}
	return -1
}