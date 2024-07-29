package main

import "fmt"

type MyHashMap struct {
    data map[int]int
}

func Constructor() MyHashMap {
    return MyHashMap{
        data: make(map[int]int),
    }
}

func (this *MyHashMap) Put(key int, value int) {
    this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
    value, exists := this.data[key]
    if !exists {
        return -1 // Key does not exist
    }
    return value
}

func (this *MyHashMap) Remove(key int) {
    delete(this.data, key)
}

func main() {
    obj := Constructor()
    obj.Put(1, 1)
    obj.Put(2, 2)
    fmt.Println(obj.Get(1)) 
    fmt.Println(obj.Get(3)) 
    obj.Put(2, 1)
    fmt.Println(obj.Get(2)) 
    obj.Remove(2)
    fmt.Println(obj.Get(2)) 
}
