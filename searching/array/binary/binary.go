package main


import "fmt"


func main(){
	a:=[12]int{6,9,17,19,26,27,39,42,47,55,69,89}
	fmt.Println(binary(a[:],39))
	
}

func binary(a []int,  target int) int{
	var start int = 0
	var end int = len(a)-1
	for start <= end{
		mid:= start +(end-start)/2
		if a[mid] > target{
			end = mid -1
		}else if a[mid] < target{
			start = mid + 1
		}else{
			return mid
		}
	}
	return -1
}