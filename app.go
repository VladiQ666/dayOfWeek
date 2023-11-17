package main

import "fmt"

func main() {
	dayOfWeek(8)
}

func dayOfWeek(day int) {
	switch day {
	case 1:
		fmt.Println("Сегодня понеднльник")
	case 2:
		fmt.Println("Сегодня вторник")
	case 3:
		fmt.Println("Сегодня среда")
	case 4:
		fmt.Println("Сегодня четверг")
	case 5:
		fmt.Println("Сегодня пятница")
	case 6:
		fmt.Println("Сегодня суббота")
	case 7:
		fmt.Println("Сегодня воскресенье")
	default:
		fmt.Println("Херню пишешь")
	}
}
