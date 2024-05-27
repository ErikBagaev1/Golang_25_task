package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Введите ширину: ")
	fmt.Scan(&num1)
	fmt.Print("Введите высоту: ")
	fmt.Scan(&num2)
	for i := 1; i <= num2; i++ {

		for j := 1; j <= num1; j++ {
			if j%2 == 0 && i%2 == 0 {
				fmt.Print(" ")

			} else if j%2 != 0 && i%2 == 0 {
				fmt.Print("*")
			}

			if j%2 == 0 && i%2 != 0 {
				fmt.Print("*")

			} else if j%2 != 0 && i%2 != 0 {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}
