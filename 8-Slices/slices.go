package main

import (
	"fmt"
	"slices"
)

func main() {
	//slices are dynamic array
	//uninitilazied slice is nil
	var nums []int

	fmt.Println(nums == nil)
	fmt.Println(len(nums))
	//initializing not nil slices
	var num = make([]int, 2, 5)
	num = append(num, 1) //appends number to the last
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	fmt.Println(num)
	fmt.Println(len(num))
	//capacity of slice
	fmt.Println(cap(num))

	//in single line

	numss := []int{}
	fmt.Println(numss)

	var num2 = make([]int, len(num))
	copy(num2, num)
	fmt.Println(num2)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice[0:3])
	fmt.Println(slice[:2])
	fmt.Println(slice[1:])
	fmt.Println(slices.Equal(slice, num))
}
