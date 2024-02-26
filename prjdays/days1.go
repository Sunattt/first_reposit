package prjdays

import (
	"fmt"
)
func PerimetrAndArea() int{
	fmt.Println("Choose your figer")
	fmt.Println("1.Периметр и площадь квадрата.")
	fmt.Println("2.Периметр и площадь треугольника.")
	fmt.Println("3.Периметр и площадь прямоугольника.")
	fmt.Println()
	var num int
	fmt.Print("--->>> ", num)
	fmt.Scan(&num)
	var value int 
	switch num {
	case 1 :
		value = squre(5)
	case 2:
		value = treugolnika(3)
	case 3:
		value = pramougolnik(4,3)
	default:
		fmt.Println("Вы выле неправильное число.")
	}
	 return value 
}
func squre(a int ) int{
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d - it is perimetr\n%d - it is area.",a*4 , a*a, )
	return a*4
}

func pramougolnik(a int, b int) int{
	fmt.Printf("%d - it is perimetr\n%d - it is area. ", (a+b)*2, a*b)
	return (a+b)*2
}

func treugolnika(a int ) int{
	fmt.Println( a*3, " - it's  perimetr")
	return a*3
}
