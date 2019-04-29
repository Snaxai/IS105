package main

import "fmt"

func main() {
	AllocateMake()
	AllocateVar()
}

//AllocateMake Her lager vi en slice uten å ha deklarert variablene på forhånd.
func AllocateMake() {
	b := make([]byte, 10, 10)
	b[3] = 11
	b[5] = 69
	b = append(b, 12)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}

//AllocateVar deklarerer vi både variabler og bruker make til å lage selve slicen
func AllocateVar() {
	b := make([]byte, 10, 100)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}
