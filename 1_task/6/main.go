package main

import (
	"fmt"
)

func main() {
	var N, K, studentNumber int

	// Запрашиваем у пользователя количество студентов
	fmt.Print("Введите количество студентов (N): ")
	fmt.Scan(&N)

	// Запрашиваем у пользователя количество групп
	fmt.Print("Введите количество групп (K): ")
	fmt.Scan(&K)

	// Запрашиваем у пользователя порядковый номер студента
	fmt.Print("Введите порядковый номер студента: ")
	fmt.Scan(&studentNumber)

	// Проверяем корректность ввода
	if N <= 0 || K <= 0 || studentNumber <= 0 || studentNumber > N {
		fmt.Println("Некорректные данные.")
		return
	}

	// Определяем, в какую группу попадет студент
	groupNumber := (studentNumber-1)%K + 1

	// Выводим результат
	fmt.Printf("Студент с номером %d попадает в группу %d.\n", studentNumber, groupNumber)
}
