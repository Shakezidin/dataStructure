package hello

import "fmt"

type GraphNode struct {
	Edges map[int][]int
}

func (g *GraphNode) BFS(value int) {
	visited := make(map[int]bool)
	visited[value] = true
	var arr = []int{value}

	for len(arr) > 0 {
		fmt.Println(arr[0])
		val := arr[0]
		arr = arr[1:]

		for _, i := range g.Edges[val] {
			if !visited[i] {
				visited[i] = true
				arr = append(arr, i)
			}
		}
	}
}

func (g *GraphNode) DFStraversa(value int) {
	fmt.Println(value)
	visited := make(map[int]bool)
	g.DFS(value, visited)
}

func (g *GraphNode) DFS(value int, visited map[int]bool) {
	arr := g.Edges[value]
	if len(arr) == 0 {
		return
	}

	visited[value] = true
	for _, val := range arr {
		if !visited[val] {
			fmt.Println(val)
			g.DFS(val, visited)
		}
	}
}

func (g *GraphNode) pathExist(a, b int) bool {
	visited := make(map[int]bool)
	visited[a] = true
	stack := []int{a}

	for len(stack) > 0 {
		val := stack[0]
		stack = stack[1:]
		if val == b {
			return true
		}
		for _, i := range g.Edges[val] {
			if !visited[i] {
				visited[i] = true
				stack = append(stack, i)
			}
		}
	}
	return false
}

func (g *GraphNode) Insert(a, b int, bi bool) {
	if bi == true {
		g.Edges[a] = append(g.Edges[a], b)
		g.Edges[b] = append(g.Edges[b], a)
	} else {
		g.Edges[a] = append(g.Edges[a], b)
	}
}

// func (g *GraphNode) Delete(a int) {
// 	visited := make(map[int]bool)
// 	delete(g.Edges, a)
// 	stack := []int{a}
// 	for len(stack) > 0 {
// 		val := stack[0]
// 		stack = stack[1:]
// 		visited[val] = true
// 		for index, value := range g.Edges[val] {
// 			if !visited[value] {
// 				stack = append(stack, value)
// 				if value == a {
// 					g.Edges[val] = append(g.Edges[val][:index], g.Edges[val][index+1:]...)
// 				}
// 			}
// 		}
// 	}
// }

func (g *GraphNode) DeleteMrthordTwo(a int) {
	for key, values := range g.Edges {
		var updatedValues []int
		for _, value := range values {
			if value != a {
				updatedValues = append(updatedValues, value)
			}
		}
		g.Edges[key] = updatedValues
	}

	delete(g.Edges, a)
}
