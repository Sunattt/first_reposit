package functions

import (
	"fmt"
	"log"
	"math/rand"
	"projects/schoole/pupil"
)

func FormVerification(student pupil.Pupil) pupil.Pupil {
	//проверка ученика директором и завучем на наличия подобающей формы
	switch student.Uniform {
	case 1:
		fmt.Println("Где ваш галстуг, молодой человек.")
		fmt.Println("Мой галстуг в моём кармане, я сейчас надену его.")
	case 2:
		fmt.Println("Почему ты пришёл на экзамен с красовками(ты что обалдел).")
		fmt.Println("Извените, я опаздывал ")
	case 0:
		fmt.Println("Отлично, хороший ученик, проходи на экзамен.")
	}
	fmt.Println("A new stage of testing begins")
	fmt.Println()
	return student
}
func ArangeSt(student pupil.Pupil) pupil.Pupil {
	//создаём очередь для ученика
	var quantity int = rand.Intn(42)
	for i := 1; i <= 41; i++ {
		fmt.Printf("Очередь %d проверка ученика перед экзаменом.\n", i)
		if quantity == i{
			fmt.Println("Пришла моя очередь.")
			fmt.Println()
			break
		}
	}
	return student
}


func TakeMark( student pupil.Pupil) pupil.Pupil{
		fmt.Println("О! Начинается проверка.")
		fmt.Println("Охраник проверяет меня на содержание шпаргалок!")
		var cheatsh bool = student.CheatThing.CheatSheets
		var phone = student.CheatThing.Phone
		if phone  && cheatsh  {
			var numbers = rand.Intn(4)
			switch numbers {
			case 2:
				fmt.Println("Дааа, как же мне повезло охранник не нашёл ни мой телефон, ни мои шпарнагалки.")
				log.Println("Есть я получил 5 спасибо телефону.")
				fmt.Println()
			case 1:
				log.Println("Блин меня выгнали с экзамена. Охранник нашёл телефон и шпаргалки./n Остаюсь на 2 год ")
				fmt.Println()
			case 0:
				fmt.Println("Оууу нет у меня отобрали телефон(за то есть шпоры), но меня оштрафовали")
				if student.Wallet.Money >= 200 {
					student.Wallet.Money -= 200
				} else if student.Wallet.Money < 200 {
					fmt.Println("Эхх жаль мне не хватает денег, но у меня есть карта :)\n Но родители узнают ")
				}
				log.Println("Ура экзамен сдан на 4")
				fmt.Println()
			case 3:
				if student.Wallet.Money >= 50 {
					student.Wallet.Money -= 50
				} else if student.Wallet.Money < 50 {
					fmt.Println("Надо было брать с собой деньги(теперь родители узнают)")
				}
				fmt.Println("Отлично главное телефон со мной, жаль что оштрафовали")
				log.Println("Какой же я молодец, телефон полезная вещь. Получил 5")
				fmt.Println()
				
			}
	} else if !phone && !cheatsh {
			fmt.Println("Я спокоен у меня нет никаких шпаргалок")
			log.Println("Экзамен сдан на 3, нормально главное не 2")
			fmt.Println()
			return student 
	} else if !phone && cheatsh {
			var numbers = rand.Intn(3)
			switch numbers {
			case 1:
				fmt.Println("Как я теперь сдам экзамен без шпаргалок(надо было спрятать в трусы).")
				if student.Wallet.Money >= 50 {
					student.Wallet.Money -= 50
				} else if student.Wallet.Money < 50 {
					fmt.Println("Надо было брать с собой деньги(теперь родители узнают)")
				}
				log.Println("Получил 2 по экзамену.")
				fmt.Println()
			case 0:
				fmt.Println("Отлично охранник не нашёл мою шпаргалку")
				log.Println("Отлично экзамен сдан на 5")
				fmt.Println()
			}
		} else if phone && !cheatsh {
			var numbers = rand.Intn(2)
			switch numbers {
			case 1:
				fmt.Println("Блин у меня нашли телефон теперь ещё и отштрафуют, я провалю экзамен ")
				if student.Wallet.Money >= 200 {
					student.Wallet.Money -= 200
				} else if student.Wallet.Money < 200 {
					fmt.Println("Эхх жаль мне не хватает денег, но у меня есть карта :)\n Но родители узнают ")
					fmt.Println()
				}
			case 2:
				fmt.Println("Охранник не нашёл мой телефон, 5 жди меня.")
				log.Println("За экзамен 5. Мне купят макбук  :)")
				fmt.Println()
			}
	}
	return student 
}

// func Subject(student pupil.Pupil) pupil.Pupil {
// 	//узнаем какой предмет будет сдавать ученик
// 	tepyExam := strudent.Exam.TypeEx
// 	if tepyExam != "устный" {
// 		sbNum := rand.Intn(3)
// 		fmt.Println("У вас письменный экзамен :) поздравляю...")
// 		switch sbNum {
// 		case 0:
// 			fmt.Println("По алгебре")
// 		case 1:
// 			fmt.Println("По литературе")
// 		case 2:
// 			fmt.Println("По физике")
// 		}
// 	} else {
// 		sbNum1 := rand.Intn(3)
// 		fmt.Println("У вас устный экзамен.")
// 		switch sbNum1 {
// 		case 1:
// 			fmt.Println("По биологии")
// 		case 0:
// 			fmt.Println("По информатике")
// 		case 2:
// 			fmt.Println("По географии")
// 		}
// 	}
// 	return student
// }
