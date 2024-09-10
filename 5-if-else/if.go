package main

import "fmt"

func main() {
	//checkAgeStatus()
	//Permission()
	declarevariableinsideif()

}

func checkAgeStatus() {
	var age int = 10

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("!Adult")
	}

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Print("Kid")
	}
}

func Permission() {

	var role = "admin"
	var hasPermission bool = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}
}

func declarevariableinsideif() {

	if a := 10; a >= 18 {
		fmt.Println("Adult")
	} else if a >= 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("kid")
	}
}


