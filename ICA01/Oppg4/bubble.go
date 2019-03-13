package main

import "fmt"

func main() {

	a := []int{3, 9, 1, 2, 8, 4, 6, 10, 7, 5}

	end := len(a) - 1

	for {

		if end == 0 {
			break
		}

		for i := 0; i < len(a)-1; i++ {

			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}

		}

		end -= 1

	}

	fmt.Println(a)

}
