package main

import (
	"fmt"
)

func isLucky(ticket int) bool {
	sumFirstHalf := (ticket/100000)%10 + (ticket/10000)%10 + (ticket/1000)%10
	sumSecondHalf := (ticket/100)%10 + (ticket/10)%10 + ticket%10
	return sumFirstHalf == sumSecondHalf
}

func main() {
	maxSize := 0
	pridI := 0
	for i := 100000; i <= 999999; i++ {

		if isLucky(i) {
			if pridI == 0 {
				pridI = i
			} else {
				distance := i - pridI
				if distance > maxSize {
					maxSize = distance
				}

				pridI = i

			}

		}

	}
	fmt.Println(maxSize)
}
