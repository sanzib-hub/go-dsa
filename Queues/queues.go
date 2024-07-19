package main

import "fmt"

// import "fmt"

type Queue struct{
	items []int
}

func(q *Queue) Enqueue(a int){
	q.items = append(q.items, a)
}

func(q *Queue) Dequeue()(int, bool){
	if len(q.items) == 0{
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue) IsEmpty() bool{
	return len(q.items) == 0
}

func (q *Queue) Peek() (int, bool){
	if len(q.items) == 0{
		return 0, false
	}
	return q.items[0], true
}

func main(){
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(7)
	queue.Enqueue(4)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Peek())

}