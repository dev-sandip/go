package main

import "fmt"

// by value
func changeNum(n int) {
	n = 5
	fmt.Println("n in changeNum:", n)
}

//by reference

func changeNumber(n *int) {
	*n = 5
	fmt.Println("n in changeNumber:", *n)
}
func main() {
	num := 1
	changeNumber(&num)
	fmt.Println("num in main:", num)
}
