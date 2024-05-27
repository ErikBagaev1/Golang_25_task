package main

import (
	"fmt"
)

func main() {
	var score int

	fmt.Println("Банкомат")

	fmt.Println("Введите сумму для снятия: ")
	fmt.Scan(&score)

	if score%100 == 0 && score <= 10000 {
		fmt.Println("Операция успешно выполнена")
	} else {
		fmt.Println("Нет средств")
	}
}
