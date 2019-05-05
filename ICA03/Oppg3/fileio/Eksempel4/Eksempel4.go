package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "Test_Eksempel4_1.txt"
	newPath := "Test_Eksempel4_2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
