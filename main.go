package main

import (
	"fmt"

	"github.com/shakezidhin/hello"
)

func main() {
	arr := []int{4, 221, 1, 0, 124, 3, 134, 13, 5521, 133, 6}
	rslt:=hello.MergeSort(arr)
	fmt.Println(rslt)
}
