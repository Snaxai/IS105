package main

import (
	"fmt"
)

func main() {
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X, %q, %b \n", i, i, i)
	}
}
