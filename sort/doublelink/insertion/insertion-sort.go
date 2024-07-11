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

func (dll *doublyLinkedList) addToList(value int) {
	newNode := &Node{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func (dll *doublyLinkedList) insertionSort() {
	if dll.head == nil {
		return
	}

	current := dll.head.next
	for current != nil {
		key := current
		prev := current.prev

		if key.next != nil {
			key.next.prev = key.prev
		} else {
			dll.tail = key.prev
		}
		if key.prev != nil {
			key.prev.next = key.next
		}

		for prev != nil && prev.value > key.value {
			prev = prev.prev
		}

		if prev == nil {
			key.next = dll.head
			key.prev = nil
			dll.head.prev = key
			dll.head = key
		} else {
			key.next = prev.next
			key.prev = prev
			if prev.next != nil {
				prev.next.prev = key
			} else {
				dll.tail = key
			}
			prev.next = key
		}

		current = current.next
	}
}

func main() {
	dll := &doublyLinkedList{}
	dll.addToList(5)
	dll.addToList(4)
	dll.addToList(3)
	dll.addToList(2)
	dll.addToList(1)
	dll.printList()
	dll.insertionSort()
	dll.printList()

}
