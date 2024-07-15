package main

import "fmt"


func main(){
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ",fibonacci(i))
		fmt.Println(" ")
	}
}
func fibonacci(i int) (ret int){
	if i ==0{
		return 0
	}
	if i==1{
		return 1
	}

	return fibonacci(i-1) + fibonacci(i-2)
}