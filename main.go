package main

import (
	"github.com/shakezidhin/hello"
)

func main() {
	grph := hello.GraphNode{Edges: make(map[int][]int)}
	grph.Insert(1,3,true)
	grph.Insert(2,3,true)
	grph.Insert(4,2,true)
	grph.Insert(4,5,true)
	grph.Insert(5,1,true)

	grph.DFS(1)
}
