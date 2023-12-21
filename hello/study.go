// package hello

// type Heapp struct {
// 	Arr []int
// }

// func Parentt(index int) int {
// 	return (index - 1) / 2
// }

// func LeftChildd(index int) int {
// 	return (2 * index) + 1
// }

// func RightChildd(index int) int {
// 	return (2 * index) + 2
// }

// func (h *Heapp) Hepify(index int) {
// 	for index >= 0 {
// 		if h.Arr[index] < h.Arr[Parentt(index)] {
// 			h.Arr[index], h.Arr[Parentt(index)] = h.Arr[Parentt(index)], h.Arr[index]
// 			index = Parentt(index)
// 		} else {
// 			break
// 		}
// 	}
// }

// func (h *Heapp) HeapifyDown(index int) {
// 	for {
// 		smallest := index
// 		left := LeftChildd(index)
// 		right := RightChildd(index)

// 		if left<len(h.Arr) && h.Arr[smallest] > h.Arr[left]{
// 			smallest = left
// 		}

// 		if right<len(h.Arr)&& h.Arr[smallest] > h.Arr[right] {
// 			smallest = right
// 		}

// 		if smallest == index {
// 			return
// 		}

// 		h.Arr[index], h.Arr[smallest] = h.Arr[smallest], h.Arr[index]
// 		index = smallest
// 	}
// }

// func (h *Heapp) Insert(val int) {
// 	h.Arr = append(h.Arr, val)
// 	h.Hepify(len(h.Arr) - 1)
// }

// func (h *Heapp) ConverHeap() {
// 	prnt := Parentt(len(h.Arr)-1)
// 	for i := prnt; i >= 0; i-- {
// 		h.HeapifyDown(i)
// 	}
// }

// func (h *Heapp) Delete() {
// 	legth := len(h.Arr) - 1
// 	h.Arr[0] = h.Arr[legth]
// 	h.HeapifyDown(0)
// }

package hello

import "fmt"

type GraphNodee struct {
	vertex map[int][]int
}

func NewGraphNodee() *GraphNodee {
	return &GraphNodee{
		vertex: make(map[int][]int),
	}
}

func (g *GraphNodee) Insert(a, b int, bi bool) {
	if bi != false {
		g.vertex[a] = append(g.vertex[a], b)
		g.vertex[b] = append(g.vertex[b], a)
	} else {
		g.vertex[a] = append(g.vertex[a], b)
	}
}

func (g *GraphNodee) BFSS(value int) {
	visited := make(map[int]bool)
	stack := []int{value}
	visited[value] = true

	for len(stack) > 0 {
		val := stack[0]
		fmt.Println(stack[0])
		stack = stack[1:]

		for _, value := range g.vertex[val] {
			if !visited[value] {
				stack = append(stack, value)
				visited[value] = true
			}
		}
	}
}

func (g *GraphNodee) DFStraversal(value int) {
	fmt.Println(value)
	visited := make(map[int]bool)
	g.DFSS(value, visited)
}

func (g *GraphNodee) DFSS(value int, visited map[int]bool) {
	arr := g.vertex[value]
	if len(arr) == 0 {
		return
	}

	visited[value] = true
	for _, val := range arr {
		if !visited[val] {
			fmt.Println(val)
			g.DFSS(val, visited)
		}
	}
}

func (g *GraphNodee) PathExistt(a, b int) bool {
	visited := make(map[int]bool)
	stack := []int{a}

	for len(stack) > 0 {
		val := stack[0]
		stack = stack[1:]

		for _, value := range g.vertex[val] {
			if !visited[value] {
				visited[value] = true
				stack = append(stack, value)
			}

			if visited[b] == true {
				return true
			}
		}
	}
	return false
}

func (g *GraphNodee) Deletee(a int) {
	for i, arr := range g.vertex {
		var arrr []int
		for j := 0; j < len(arr); j++ {
			if arr[j] != a {
				arrr = append(arrr, arr[j])
			}
		}
		g.vertex[i] = arrr
	}

	delete(g.vertex, a)
}
