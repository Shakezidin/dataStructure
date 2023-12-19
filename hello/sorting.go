package hello

func Bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
	return arr
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		smallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		arr[i], arr[smallest] = arr[smallest], arr[i]
	}
	return arr
}

func Sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func QuickSort(arr []int) {
	quickSorthelper(arr, 0, len(arr)-1)
}

func quickSorthelper(arr []int, strt, end int) {
	if strt >= end {
		return
	}

	i := strt + 1
	j := end
	pivot := strt

	for i <= j {
		if arr[i] > arr[pivot] && arr[j] < arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
			i++
		} else if arr[i] < arr[pivot] {
			i++
		} else if arr[j] > arr[pivot] {
			j--
		}
	}
	arr[i-1], arr[pivot] = arr[pivot], arr[i-1]
	quickSorthelper(arr, strt, j-1)
	quickSorthelper(arr, j+1, end)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return MergeSortHelper(left, right)
}

func MergeSortHelper(left, right []int) []int {
	arr := make([]int, len(left)+len(right))
	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			k++
			i++
		} else {
			arr[k] = right[j]
			j++
			k++
		}
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
	return arr
}

