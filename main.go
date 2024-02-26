package main

import (
	"fmt"
	"log"
	"projects/prjdays"
	"projects/schoole"
)

func main() {
	log.Println("Hi, choose your day :) ")
	var days int
	fmt.Scan(&days)
	switch days {
	case 1:
		fmt.Println(prjdays.PerimetrAndArea())
	case 2:
		fmt.Println(fmt.Println("До юбилея вам нужно ещё", prjdays.Anniversary(), "лет."))
	case 3:
		fmt.Println("Нам ничего не задавали")
	case 4:
		fmt.Println("Бонусные задания")
		fmt.Println(prjdays.Bonus(""))
		fmt.Println()
		fmt.Println(prjdays.Fibonachi(""))
	case 5:
		fmt.Println(prjdays.Day5())
	default:
		fmt.Println("Ты написал неправильное число.")
	
	}
	log.Println("The programm game over")
}
