package main

import (
	"fmt"
	"strings"
)

func main() {
	var day string
	fmt.Print("Введите будний день недели: пн, вт, ср, чт, пт:s")
	fmt.Scan(&day)

	day = strings.ToLower(day)

	switch day {
	case "пн":
		fmt.Println("Это следующие рабочие дни: пн, вт, ср, чт, пт")
	case "вт":
		fmt.Println("Это следующие рабочие дни: вт, ср, чт, пт")
	case "ср":
		fmt.Println("Это следующие рабочие дни: ср, чт, пт")
	case "чт":
		fmt.Println("Это следующие рабочие дни: чт, пт")
	case "пт":
		fmt.Println("Это следующие рабочие дни: пт")
	default:
		fmt.Println("Некорректное название дня недели")
	}
}
