package main

import "fmt"

func main() {

	m := make(map[string]string)
	m["name"] = "Sandip"
	m["address"] = "Dang"

	fmt.Println(m["name"], m["address"])

	//If key doesnot exits in the map it returns zero value ie: empty  string if string , zero if integer and false for boolean
}
