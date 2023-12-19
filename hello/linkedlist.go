package hello

import "fmt"

func TwoSum(arr []int, number int) []int {
	nums := make(map[int]int)
	result := []int{}

	for i := 0; i < len(arr); i++ {
		val := arr[i] - number
		if _, exists := nums[val]; exists {
			result = append(result, val, arr[i])
			return result
		}
		nums[arr[i]] = i
	}
	return result
}

func MoveSide(arr []int, nums int) []int {
	k := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[i] == nums {
			for k >= 0 {
				if arr[k] != nums {
					temp := arr[k]
					arr[k] = arr[i]
					arr[i] = temp
					break
				} else {
					k--
				}
			}
		}
	}
	return arr
}

func MoveSideTwo(arr []int, nums int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left] == nums {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		} else {
			left++
		}
	}

	return arr
}

type node struct {
	Data int
	Next *node
}

type LinkedList struct {
	Head *node
	Tail *node
}

func (l *LinkedList) InsertEnd(val int) {
	newnode := node{
		Data: val,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = &newnode
		l.Tail = &newnode
		return
	} else {
		l.Tail.Next = &newnode
		l.Tail = &newnode
		return
	}
}

func (l *LinkedList) InserBegenning(val int) {
	newnode := node{
		Data: val,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = &newnode
		l.Tail = &newnode
	} else {
		newnode.Next = l.Head
		l.Head = &newnode
	}
}

func (l *LinkedList) InsertNextValue(val, data int) {
	newnode := node{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = &newnode
		l.Tail = &newnode
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Data == val {
			newnode.Next = current.Next
			current.Next = &newnode
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Deletenode(data int) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == data {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return
	}

	prev := l.Head
	current := l.Head.Next

	for current != nil {
		if current.Data == data {
			prev.Next = current.Next
			if current.Next == nil {
				l.Tail = prev
			}
			return
		}
		current = current.Next
		prev = prev.Next
	}
}

func (l *LinkedList) DeleteDuplicates() {
	if l.Head == nil {
		return
	}

	current := l.Head

	for current != nil {
		runner := current
		for runner.Next != nil {
			if runner.Next.Data == current.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		current = current.Next
	}
}

func (l *LinkedList) Display() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

type Stack struct {
	top *node
}

func (l *Stack) Push(val int) {
	newnode := node{
		Data: val,
		Next: nil,
	}
	if l.top == nil {
		l.top = &newnode
		return
	} else {
		newnode.Next = l.top
		l.top = &newnode
		return
	}
}

func (l *Stack) Pop() {
	if l.top == nil {
		fmt.Println("empty")
		return
	} else {
		l.top = l.top.Next
		return
	}
}

func addNumberToString(inputString string, number int) string {
	result := ""

	for _, char := range inputString {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			newChar := (char+rune(number)-'a')%26 + 'a'
			result += string(newChar)
		} else {
			result += string(char)
		}
	}

	return result
}

func AddChar(str string, num int) string {
	newStr := ""
	newNum := num % 26
	for _, char := range str {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			newChar := char + rune(newNum)
			newStr += string(newChar)
		} else {
			newStr += string(char)
		}
	}
	return newStr
}


