package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4}
	userData := map[string]string{"first-name": "Sandip", "last-name": "Sapkota", "profession": "Backend Developer"}
	// classicalIteration(numbers)
	// usingRange(numbers)
	rangeInMap(userData)
	//unicode code
	//starting byte of rune
	for i, v := range "golang" {
		fmt.Println(i, v)
	}
}

func classicalIteration(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func usingRange(numbers []int) {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
		fmt.Println(number)
	}
	fmt.Println("The sum of the array of numbers is ", sum)
}
func rangeInMap(data map[string]string) {
	for key, value := range data {
		fmt.Println(key, value)
	}
	//If single variable given then it will return keys
}
