package main

import (
	"fmt"
)

func main() {
	passengers := [25]int{0, 0, 0, 3, 0, 0, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	reis := 0

	for {
		allZero := true
		for _, p := range passengers {
			if p > 0 {
				allZero = false
				break
			}
		}

		if allZero {
			break
		}

		for i := len(passengers) - 1; i > 0; i-- {
			for passengers[i] > 0 {
				vagon := 0

				if passengers[i] >= 2 {
					vagon = 2
					passengers[i] -= 2
				} else {
					vagon = passengers[i]
					passengers[i] = 0
				}

				if vagon == 1 {
					for j := i - 1; j > 0; j-- {
						if passengers[j] > 0 {
							vagon = 2
							passengers[j]--
							break
						}
					}
				}

				reis++
				fmt.Printf("Лифт забирает %d пассажиров с %d этажа, всего рейсов: %d\n", vagon, i, reis)
			}
		}
	}

	fmt.Printf("Всего рейсов: %d\n", reis)
}
