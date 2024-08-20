package main

import "fmt"


type Tree struct{
	val int
	left *Tree
	right *Tree
}


func(tree *Tree) Insert(value int) {
	if tree == nil{
		return
	} else if value < tree.val{
		if tree.left == nil{
			tree.left = &Tree{val: value}
		} else{
			tree.left.Insert(value)
		}
	}else{
		if tree.right == nil{
			tree.right = &Tree{val: value}
		} else{
			tree.right.Insert(value)
		}
	}
}

func(tree *Tree) InOrderTraversal(){
	if tree == nil{
		return
	}
	tree.left.InOrderTraversal()
    fmt.Printf("%d ", tree.val)
	tree.right.InOrderTraversal()
}
func main() {
    root := &Tree{val: 10}
    root.Insert(5)
    root.Insert(15)
    root.Insert(3)
    root.Insert(7)
    root.Insert(12)
    root.Insert(18)

    fmt.Print("InOrder Traversal: ")
    root.InOrderTraversal() 
}