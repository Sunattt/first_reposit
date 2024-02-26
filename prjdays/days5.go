package prjdays

import (
	"projects/prjdays/customer"
	"projects/prjdays/operations"
	"log"
	"math/rand"
)

func Day5() {
	//Тестовый проект - "Дурная больница"
	//Авторы - студенты HumoLab
	var Customer *customer.Patient = &customer.Patient{
		//rand.Intn(126) + 5 означает, что мы хотим рандомное число от 5 до 130. Т.е от 0 до 126 исключительно, и к итоговому результату всегда дописывается + 5
		Name: "Гордон Фриманович Ломов", Age: rand.Intn(126) + 5, Order: 0, Organs: customer.Organs{Leg: 2, Hands: 2, Livers: 2, Heart: false}}
	log.Println("Я зашёл забрать чек")
	Customer = operations.TakeOrder(Customer)
	log.Println("Я забрал чек. \nМой номер:", Customer.Order)
	log.Println("Я пошёл на операцию.")
	Customer = operations.BeginOperation(Customer)
	//Проверка результата после операции
	if !Customer.Organs.Heart {
		log.Printf("Я Вернулся! Но что-то со мной не так: %+v", Customer)
	} else {
		log.Printf("Я Вернулся! И чувствую себя прекрасно!: %+v", Customer)
	}
}
