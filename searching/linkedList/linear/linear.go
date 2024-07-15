package main

import "fmt"


type Node struct{
	Val int
	next *Node
}


func main(){
	node1:=&Node{Val: 5}
	node2:=&Node{Val: 9}
	node3:=&Node{Val: 24}
	node4:=&Node{Val: 56}
	node5:=&Node{Val: 2}
	node6:=&Node{Val: 77}
	node7:=&Node{Val: 17}
	
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7

	// PrintList(node1)
	fmt.Println(linear(node1,77))

}

func linear(head *Node, target int) bool{
	if head == nil{
		return false
	}
	for current:= head; current !=nil; current = current.next{
		if current.Val == target{
			return true
		}
	}
	return false
}

func PrintList(head *Node){
	current := head

	for current !=nil{
		fmt.Printf("%d-->", current.Val)
		current = current.next
	}
	fmt.Println("nil")
}