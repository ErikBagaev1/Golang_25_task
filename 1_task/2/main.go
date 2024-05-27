package main

import (
	"fmt"
)

func main() {
	var score1, score2, score3 int

	fmt.Println("Три числа")

	fmt.Println("Введите первое число: ")
	fmt.Scan(&score1)

	fmt.Println("Введите второе число: ")
	fmt.Scan(&score2)

	fmt.Println("Введите третье число: ")
	fmt.Scan(&score3)

	if score1 > 5 || score2 > 5 || score3 > 5 {
		fmt.Println("Среди введеных чисел есть число больше 5")
	} else {
		fmt.Println("Среди введеных чисел нет числа больше 5")
	}
}
