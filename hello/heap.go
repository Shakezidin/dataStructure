package hello

func LeftChild(index int) int {
	return (2 * index) + 1
}

func RightChild(index int) int {
	return (2 * index) + 2
}

func Parent(index int) int {
	if index == 0 {
		return 0
	}
	return (index - 1) / 2
}

type Heap struct {
	arr []int
}

func (h *Heap) HeapifyDown(index int) {
	for {
		smallest := index
		left := LeftChild(smallest)
		right := RightChild(smallest)

		if h.arr[smallest] < h.arr[left] && len(h.arr) < left {
			smallest = left
		}
		if h.arr[smallest] < h.arr[right] && len(h.arr) < right {
			smallest = right
		}

		if smallest == index {
			return
		}

		h.arr[smallest], h.arr[index] = h.arr[index], h.arr[smallest]
		index = smallest
	}
}

func (h *Heap) heapify(index int) {
	for index >= 0 {
		parent := Parent(index)
		if h.arr[parent] < h.arr[index] {
			h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent]
			index = parent
		} else {
			return
		}
	}
}

func (h *Heap) Insert(data int) {
	h.arr = append(h.arr, data)
	index := len(h.arr) - 1
	h.heapify(index)
}

func (h *Heap) Delete() {
	length := len(h.arr) - 1
	h.arr[0] = h.arr[length]
	h.HeapifyDown(0)
}

func (h *Heap)BuildHeap(){
	for i:=len(h.arr)-1;i>=0;i--{
		h.HeapifyDown(i)
	}
}