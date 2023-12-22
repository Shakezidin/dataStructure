package hello

import "fmt"

type GraphNodee struct {
	edge map[int][]int
}

func (g *GraphNodee) Insert(a, b int, bi bool) {
	if bi != false {
		g.edge[a] = append(g.edge[a], b)
		g.edge[b] = append(g.edge[b], a)
	} else {
		g.edge[a] = append(g.edge[a], b)
	}
}

func (g *GraphNodee) BFS(val int) {
	visited := make(map[int]bool)
	Stack := []int{val}
	visited[val] = true

	for len(Stack) > 0 {
		val := Stack[0]
		Stack = Stack[1:]
		fmt.Println(val)

		for _, value := range g.edge[val] {
			if !visited[value] {
				visited[value] = true
				Stack = append(Stack, value)
			}
		}
	}
}

func (g *GraphNodee) DFS(val int) {
	visited := make(map[int]bool)
	fmt.Println(val)
	g.DFSHelper(val, visited)
}

func (g *GraphNodee) DFSHelper(val int, visited map[int]bool) {
	arr := g.edge[val]
	if len(arr) == 0 {
		return
	}
	visited[val] = true
	for _, i := range g.edge[val] {
		if !visited[i] {
			fmt.Println(i)
			visited[i] = true
			g.DFSHelper(i, visited)
		}
	}
}
