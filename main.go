package main

import (
	"fmt"

	"github.com/shakezidhin/hello"
)

func main() {
	graph := hello.NewGraphNodee()

	graph.Insert(1, 3, true)
	graph.Insert(2, 4, false)
	graph.Insert(3, 5, true)
	graph.Insert(1, 2, true)
	graph.Insert(4, 5, false)

	graph.DFStraversal(1)

	fmt.Println("path exist",graph.PathExistt(2,7))
	graph.Deletee(1)
	fmt.Println("path exist two",graph.PathExistt(2,7))
	graph.BFSS(2)
}
