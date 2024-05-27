package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Println(i)
	}

}
