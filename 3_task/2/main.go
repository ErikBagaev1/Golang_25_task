package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Введите число 1: ")
	fmt.Scan(&num1)
	fmt.Print("Введите число 2: ")
	fmt.Scan(&num2)
	for i := num1; i <= num1+num2; i++ {
		fmt.Println(i)
	}

}
