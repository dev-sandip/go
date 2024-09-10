package main

import "fmt"

func main() {
	// var nums [4]int
	// nums[0] = 1
	// fmt.Println(len(nums), nums[0])
	// fmt.Println(nums)

	// var vals [3]bool
	// fmt.Println(vals)

	// var name [4]string
	// name[0] = "Sandip"
	// fmt.Println(name)
	//declaring in single line
	num := [4]int{1, 2, 3}
	fmt.Println(num)

	//two dimnesional array

	number := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(number)

}
