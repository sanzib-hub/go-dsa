package main


import "fmt"


func main(){
	a := [10]int{2, 12, 15, 11, 7, 19, 45, 85,9, 23}
	// b :=[10]int{}
	// for i := 0; i < len(a); i++ {
	// 	b[i] = a[i]-1
	// }
	// fmt.Println(b)

	fmt.Println(linear(a[:], 7))
}

func linear(a []int, target int) int{
	for i := 0; i < len(a); i++ {
		if a[i] == target{
			return i
		}
	}
	return -1
}