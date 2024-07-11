package main

import (
	"fmt"
	// "go/build/constraint"
)

type Node struct {
	value int
	next  *Node
}

func main() {
	node1 := &Node{value: 5}
	node2 := &Node{value: 4}
	node3 := &Node{value: 3}
	node4 := &Node{value: 2}
	node5 := &Node{value: 1}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	fmt.Println("Before Sort")
	printList(node1)

	c := insertionSort(node1)

	fmt.Println("After Sort")
	printList(c)

}

func printList(head *Node) {

	current := head
	for current != nil {
		fmt.Printf("%d --->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func insertionSort(head *Node) *Node {
	if head == nil {
		return nil
	}

	sorted := &Node{value: head.value}
	current := head.next
	for current != nil {
		next := current.next
		sorted = insertSortedlist(sorted, current)
		current = next
	}
	return sorted
}

func insertSortedlist(sorted *Node, newNode *Node) *Node {
	if sorted == nil || newNode.value <= sorted.value {
		newNode.next = sorted
		return newNode
	}

	current := sorted

	for current.next != nil && current.next.value < newNode.value {
		current = current.next

	}
	newNode.next = current.next
	current.next = newNode

	return sorted
}
