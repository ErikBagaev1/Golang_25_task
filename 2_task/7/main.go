package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	var task string

	temp := 5
	for {
		if num == temp {
			fmt.Println("Ваше число: ", num)
			break
		} else {
			fmt.Printf("Ваше число больше? %d (Да, Нет)", temp)
			fmt.Scan(&task)
			if task == "Да" {
				temp += (temp) / 2
			} else {
				temp = (temp) / 2
			}
		}
	}

}
