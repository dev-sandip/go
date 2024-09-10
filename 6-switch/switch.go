package main

import (
	"fmt"
	"time"
)

func main() {
	// basicSwitch()
	// multipleConditionSwitch()
	typeSwitch()

}

func basicSwitch() {
	var day string = "Sunday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is a different day")
	}
}
func multipleConditionSwitch() {

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work Day")
	}
}

func typeSwitch() {
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer!")
		case string:
			fmt.Println("It's an string!")
		case bool:
			fmt.Println("It's an boolean")
		default:
			fmt.Println("Other", t)

		}

	}

	whoAmI("")
}
