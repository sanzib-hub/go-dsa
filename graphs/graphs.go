package main

import "fmt"

type Graph struct{
	vert map[int][]int
}

func(g *Graph) addEdge(v,w int){
	if g.vert  == nil{
		g.vert = make(map[int][]int)
	}
	g.vert[v] = append(g.vert[v], w)
    g.vert[w] = append(g.vert[w], v) 
}
func (g *Graph) DFS(start int, visited map[int]bool) {
    visited[start] = true
    fmt.Println(start)

    for _, v := range g.vert[start] {
        if !visited[v] {
            g.DFS(v, visited)
        }
    }
}
func (g *Graph) BFS(start int) {
    visited := make(map[int]bool)
    queue := []int{start}

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]

        if !visited[vertex] {
            fmt.Println(vertex)
            visited[vertex] = true

            for _, v := range g.vert[vertex] {
                if !visited[v] {
                    queue = append(queue, v)
                }
            }
        }
    }
}

func main(){
	g := Graph{}
    g.addEdge(0, 1)
    g.addEdge(0, 2)
    g.addEdge(1, 2)
    g.addEdge(2, 0)
    g.addEdge(2, 3)
    g.addEdge(3, 3)

    fmt.Println("Depth-First Search (starting from vertex 2):")
    g.DFS(2, make(map[int]bool))

    fmt.Println("\nBreadth-First Search (starting from vertex 2):")
    g.BFS(2)
	

}