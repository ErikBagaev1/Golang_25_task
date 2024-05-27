package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	fmt.Print("Введите третье число: ")
	fmt.Scan(&num3)

	if num1 > 0 || num2 > 0 || num3 > 0 {
		fmt.Println("Хотя бы одно из введённых чисел положительное.")
	} else {
		fmt.Println("Ни одно из введённых чисел не является положительным.")
	}
}
