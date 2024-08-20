package main

import "fmt"

type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

// Constructor initializes the Trie.
func Constructor() Trie {
    return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert adds a word into the Trie.
func (this *Trie) Insert(word string) {
    node := this.root
    for _, ch := range word {
        if _, exists := node.children[ch]; !exists {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

// Search checks if the word is in the Trie.
func (this *Trie) Search(word string) bool {
    node := this.root
    for _, ch := range word {
        if _, exists := node.children[ch]; !exists {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
    node := this.root
    for _, ch := range prefix {
        if _, exists := node.children[ch]; !exists {
            return false
        }
        node = node.children[ch]
    }
    return true
}

func main() {
    obj := Constructor()
    obj.Insert("apple")
    fmt.Println(obj.Search("apple"))   // returns true
    fmt.Println(obj.Search("app"))     // returns false
    fmt.Println(obj.StartsWith("app")) // returns true
    obj.Insert("app")
    fmt.Println(obj.Search("app"))     // returns true
}
