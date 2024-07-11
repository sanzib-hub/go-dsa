package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type doublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *doublyLinkedList) printList() {
	current := dll.head

	for current != nil {
		fmt.Printf("%d----->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (dll *doublyLinkedList) printListReverse() {
	current := dll.tail

	for current != nil {
		fmt.Printf("%d--->", current.value)
		current = current.prev

	}
	fmt.Println("nil")
}

func(dll *doublyLinkedList) addToEnd(value int){
	newNode := &Node{value: value}

	if dll.head == nil{
		dll.head = newNode
		dll.tail = newNode
	}else{
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func(dll *doublyLinkedList) bubbleSort(){
	if dll.head == nil{
		return
	}
	var swapped bool
	for {
		swapped = false
		current := dll.head
		for current != nil && current.next !=nil{
			if current.value > current.next.value{
				current.value, current.next.value = current.next.value, current.value
				swapped = true
			}
			current = current.next
		}
		if !swapped{
			break
		}
	}
}

func main(){
	dll := &doublyLinkedList{}
	dll.addToEnd(5)
	dll.addToEnd(4)
	dll.addToEnd(3)
	dll.addToEnd(2)
	dll.addToEnd(1)

	fmt.Println("Print List")
	dll.printList()

	dll.bubbleSort()

    fmt.Println("Sorted List")
    dll.printList()
	
// 	fmt.Println("Print list in reverse")
// 	dll.printListReverse()
}