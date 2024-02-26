package prjdays

import (
	"fmt"
	"math"
)
func Anniversary() int {
	fmt.Println("Hello, It's program for finding how\nmany days are left until the anniversary.")
	fmt.Println("How old are you?")
	var age int
	fmt.Print("I'm ")
	fmt.Scan(&age)
	days := 10 - int(float64(age)*0.1 - math.Floor(float64(age) * 0.1)) // remainder
	return days
}