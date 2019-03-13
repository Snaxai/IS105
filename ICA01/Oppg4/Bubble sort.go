package main

import (
	"fmt"
)

var toBeSorted [10]int = [10]int{1, 3, 2, 9, 4, 8, 6, 7, 5, 0}

func bubbleSort(input [10]int) {

	n := 10

	swapped := true

	for swapped {

		swapped = false

		for i := 1; i < n; i++ {

			if input[i-1] > input[i-1] {

				fmt.Println("Swapping")

				input[i], input[i-1] = input[i-1], input[i]

				swapped = true
			}
		}
	}

	fmt.Println(input)
}
func main() {
	fmt.Println("Hei på deg")
	bubbleSort(toBeSorted)
}
