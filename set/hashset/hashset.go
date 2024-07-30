package main

import "fmt"

type HashSet struct {
	data map[int]struct{}
}

func NewHashSet() HashSet {
	return HashSet{
		data: make(map[int]struct{}),
	}
}


func (s *HashSet) Add(key int) {
	s.data[key] = struct{}{}
}


func (s *HashSet) Remove(key int) {
	delete(s.data, key)
}

func (s *HashSet) Contains(key int) bool {
	_, exists := s.data[key]
	return exists
}

func main() {
	set := NewHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(4)) 

	set.Remove(2)
	fmt.Println(set.Contains(2)) 
}
