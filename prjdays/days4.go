package prjdays

import (
	"fmt"
)
func Bonus(a string) string{
	for i := 1; i <= 9; i++{
		for a := 1; a < 10; a++{
			fmt.Printf("%d * %d = %d  ", a, i, i*a)
		}
		fmt.Println()
	} 
	return a
}

func Fibonachi(a string) string{
	// var numbers int
	// fmt.Scan(&numbers)
	fb1 := 1
	fn2 := 1
	fb_sum := 0 
	for i := 0; i >= 10; i++{
		fb_sum = fb1 + fn2
		fb1 = fn2
		fn2 = fb_sum
        fmt.Printf("%v, %v, ", fb1, fn2)
	}
	return a
}