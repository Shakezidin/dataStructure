package hello

func binary(arr []int, val int) bool {

	frst := 0
	end := len(arr) - 1

	for frst < end {
		middle := (end + frst) / 2

		if arr[middle] == val {
			return true
		} else if val > arr[middle] {
			frst = middle
		} else if val < arr[middle] {
			end = middle
		}
	}
	return false
}
