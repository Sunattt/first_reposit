package prjdays

import (
	"fmt"
	"math"
)

func Anniversary() float64{
	fmt.Println("Hello, It's program for finding how\nmany days are left until the anniversary.")
	fmt.Println("How old are you?")
	var age float64
	fmt.Print("I'm ")
	fmt.Scan(&age)
	//days := 10 - (age % 10)
	fmt.Println(math.Floor(age*0.1))
	days := 10 - (age - math.Floor(age*0.1) * 10) // remainder
	return days
}
