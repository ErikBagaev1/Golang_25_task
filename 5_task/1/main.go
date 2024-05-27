package main

import (
	"fmt"
	"strings"
)

func main() {
	var month string
	fmt.Print("Введите название месяца: ")
	fmt.Scan(&month)

	month = strings.ToLower(month)

	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Это зима")
	case "март", "апрель", "май":
		fmt.Println("Это весна")
	case "июнь", "июль", "август":
		fmt.Println("Это лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Это осень")
	default:
		fmt.Println("Некорректное название месяца")
	}
}
