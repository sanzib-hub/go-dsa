package main


import (
	"fmt"
)


func main(){
	a:=[5]int{5,4,3,2,1}
	b:=[]int{}
	b=append(b, a[:]...)
	fmt.Println("Before sort",b)

	for i:=1;i<len(b);i++{
		index:=b[i]
		j:=i-1
		for j>=0 && b[j] >index{
			b[j+1]=b[j]
			j=j-1
		}
		b[j+1] = index
	}
	fmt.Println(b)
}