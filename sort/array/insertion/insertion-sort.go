package main

import (
	"fmt"
)

func main(){
	a:=[5]int{5,4,3,2,1}
	b:=[]int{}

	b=append(b,a[:]...)
	 fmt.Println("Before sort",b)
	for i := 1; i < len(b)-1; i++ {
		key :=b[i]
		j:=i-1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j=j-1
		}
		a[j+1] = key
	}
	fmt.Println("After sort",b)
}