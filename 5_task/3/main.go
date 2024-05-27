package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lemonadeChange(bills []int) bool {
	fiveCount, tenCount := 0, 0

	for _, bill := range bills {
		switch bill {
		case 5:
			fiveCount++
		case 10:
			if fiveCount == 0 {
				return false
			}
			fiveCount--
			tenCount++
		case 20:
			if tenCount > 0 && fiveCount > 0 {
				tenCount--
				fiveCount--
			} else if fiveCount >= 3 {
				fiveCount -= 3
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Введите купюры через запятую: ")
	fmt.Scanln(&input)

	strBills := strings.Split(input, ",")
	bills := make([]int, len(strBills))

	for i, strBill := range strBills {
		bill, err := strconv.Atoi(strBill)
		if err != nil {
			fmt.Println("Ошибка: введите только числа.")
			return
		}
		bills[i] = bill
	}

	result := lemonadeChange(bills)
	fmt.Println(result)
}
