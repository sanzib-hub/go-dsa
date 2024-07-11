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
		fmt.Printf("%d--->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func(dll *doublyLinkedList) addToList(value int){
	newNode := &Node{value: value}
	if dll.head == nil{
		dll.head = newNode
		dll.tail = newNode
	} else{
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func (dll *doublyLinkedList) selectionSort(){
	if dll.head == nil{
		return
	}
	for current:= dll.head; current != nil; current=current.next {
		min:= current
		for walker:=current.next; walker !=nil; walker=walker.next{
			if walker.value < min.value{
				min = walker
			}
		}
		current.value, min.value = min.value, current.value
	}
}

func main(){
	dll := &doublyLinkedList{}
	dll.addToList(5)
	dll.addToList(4)
	dll.addToList(3)
	dll.addToList(2)
	dll.addToList(1)
	dll.printList()
	dll.selectionSort()
	dll.printList()
}