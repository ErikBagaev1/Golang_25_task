package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Введите координату X: ")
	fmt.Scan(&x)

	fmt.Print("Введите координату Y: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка находится в первой четверти.")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка находится во второй четверти.")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка находится в третьей четверти.")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка находится в четвёртой четверти.")
	} else if x == 0 && y != 0 {
		fmt.Println("Точка находится на оси Y.")
	} else if y == 0 && x != 0 {
		fmt.Println("Точка находится на оси X.")
	} else {
		fmt.Println("Точка находится в начале координат.")
	}
}
