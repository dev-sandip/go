package main

import "fmt"

func main() {
	fmt.Println(add(500, 67))
}
// add takes two integers a and b, and returns their sum.
func add(a int, b int) int {
	return a + b
}
