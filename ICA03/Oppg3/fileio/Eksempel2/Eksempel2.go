//Korter ned en bestemt fil.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Truncate a file to 100 bytes. If file
	// is less than 100 bytes the original contents will remain
	// at the beginning, and the rest of the space is
	// filled will null bytes. If it is over 100 bytes,
	// Everything past 100 bytes will be lost. Either way
	// we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely empty file

	err := os.Truncate("Test_Eksempel2.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

	filstorrelse()
}

//Lagde en ny funksjon som skriver ut filstørrelsen for at vi kan se at programmet faktisk fungerer som den skal.
//Altså at den korter filen ned til 100 bytes.
func filstorrelse() {
	size, err := ioutil.ReadFile("Test_Eksempel2.txt")

	if err != nil {
		log.Panicf("failed at reading data from file: %s", err)
	}

	fmt.Printf("Filstørrelse: %d bytes\n", len(size))
}
