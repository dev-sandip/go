package main

import "fmt"

func main() {
	var a int = 1
	var b float32 = 77.77

	var isAdult bool = true

	var name = "Sandip Sapkota"

	age := 22
	fmt.Println(a)
	fmt.Println(b)
	if isAdult {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not Adult")

	}

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
}
