package main

import "fmt"

func main() {

	// goLoop()
	// classicalLoop()
	Range()
}

// classical for loop

func goLoop() {
	i := 0
	for i <= 3 {
		i = i + 1
		fmt.Println(i)

	}
}

func classicalLoop() {

	for i := 1; i <= 3; i++ {

		if i == 2 {
			continue
		}
		fmt.Println("Hello World", i)
	}
}

//using range

func Range() {
	for i := range 3 {
		fmt.Println(i)
	}
}
