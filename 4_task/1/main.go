package main

import (
	"fmt"
	"strconv"
)

func isMirrorNumber(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	for num := 100000; num <= 999999; num++ {
		if isMirrorNumber(num) {
			count++
		}
	}
	fmt.Println("Количество зеркальных билетов:", count)
}
