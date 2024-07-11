package main

import "fmt"


type Node struct{
	value int
	next *Node
}

func main(){
	node1:= &Node{value: 5}
	node2:= &Node{value: 4}
	node3:= &Node{value: 3}
	node4:= &Node{value: 2}
	node5:= &Node{value: 1}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	fmt.Println("Before sort")
	printList(node1)

}

func printList(head *Node){
	current := head

	for current != nil{
		fmt.Printf("%d-->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}