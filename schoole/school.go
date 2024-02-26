package schoole

import (
	"log"
	"fmt"
	"math/rand"
	"projects/schoole/functions"
	"projects/schoole/pupil"
)

func MiniSchoole(){
	//начало программы
	//автор Сунат
	log.Println("Hi, It's mini project the School exame.")
	// узнаем тип экзамена письменный или устный
	a := rand.Intn(2)
	var b string
	if a == 1 {
	    b = "устный" 
	}else if a == 0 {
		b = "письменный"
	}
	
// узнаем какой экзамен будет сдавать ученик 

	var name string
	if b == "устный" {
		sbNum := rand.Intn(2)
	
			switch sbNum {
				case 0:
					name = "По алгебре"
				case 1:
					name = "По литературе"					
				}
	} else {
		sbNum1 := rand.Intn(4)
		switch sbNum1 {
		case 1:
			name = "По биологии"
		case 0:
			name = "По информатике"
		case 2:
			name = "По географии"
		case 3:
			name = "По физике"
		}
	}
	
		var Scholar pupil.Pupil = pupil.Pupil{
		Name: "Тимофей Александрович Швили", Age: rand.Intn(13) + 5, Class: rand.Intn(4) + 7,
		Wallet:     pupil.Cash{Money: rand.Intn(100), Card: true},
		CheatThing: pupil.CheatThing{CheatSheets: false, Phone: true}, Uniform: rand.Intn(2),
		Exam: pupil.TpExam{TypeEx: b, WhichEx: name}}
fmt.Printf("%+v", Scholar)
fmt.Println()
fmt.Println()
    Scholar = functions.FormVerification(Scholar)  

	//передмет
	log.Printf("Надо бы узнать, какой предмет мы будем сдавать\n Мы сдаём %v экзамен %v", b, name)
	//проврека директора на форму
	log.Println(Scholar)
	fmt.Println()
	//Цикл, очердь3
	log.Println(functions.ArangeSt(Scholar))
	fmt.Println()
	//Проверка охранника
	log.Println(functions.TakeMark(Scholar))
	fmt.Println()

fmt.Printf("%+v",Scholar)


}