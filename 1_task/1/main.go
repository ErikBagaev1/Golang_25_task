package main

import (
	"fmt"
)

func main() {
	var score1, score2, score3 int

	fmt.Println("Баллы ЕГЭ.")

	fmt.Print("Баллы за первый экзамен: ")
	fmt.Scan(&score1)

	fmt.Print("Баллы за второй экзамен: ")
	fmt.Scan(&score2)

	fmt.Print("Баллы за третий экзамен: ")
	fmt.Scan(&score3)

	total := score1 + score2 + score3
	fmt.Println("Сумма проходных баллов: 275")
	fmt.Println("Количество набранных баллов: ", total)
	if total >= 275 {
		fmt.Println("Поздравляем! Вы поступили в институт.")
	} else {
		fmt.Println("К сожалению, вы не поступили в институт.")
	}
}
