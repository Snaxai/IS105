package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filinfo()
}

func filinfo() {
	fmt.Printf("\nInformasjon om filen text.txt: \n\n")

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("Kunne ikke åpne filen: text.txt %s\n", err)
		os.Exit(1)
	}

	//Skriver ut størrelsen på text.txt filen som vi opprettet. Først blir størrelsen skrevet ut i bytes, deretter bruker vi denne
	//størrelsen til å skrive ut verdien i kibibytes, mibibytes og gibibytes.

	size, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Panicf("failed at reading data from file: %s", err)
	}
	fmt.Printf("Filstørrelse: %d bytes\n", len(size))
	i := len(size)
	f := float64(i)
	fmt.Printf("Filstørrelse: %f Kibibytes\n", f/1e3)   //Tar utgangspunkt i 190 og flytter kommaet tre plasser til venstre vha. "f/1e3". For å skrive ut i kibibytes.
	fmt.Printf("Filstørrelse: %f Mibibytes\n", f/1e6)   //Tar utgangspunkt i 190 og flytter kommaet seks plasser til venstre vha. "f/1e6". For å skrive ut i mibibytes.
	fmt.Printf("Filstørrelse: %f Gibigytes\n\n", f/1e9) //Tar utgangspunkt i 190 og flytter kommaet ni plasser til venstre vha. "f/1e9". For å skrive ut i gibibytes.

	//Skriver ut stats til filen text.txt
	stats, err := file.Stat()
	if err != nil {
		fmt.Printf("Kunne ikke finne stats: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Directory: %t\n", stats.IsDir())
	fmt.Printf("Regular: %t\n", stats.Mode().IsRegular())
	fmt.Printf("Unix permissions: %s\n", stats.Mode().String())
	fmt.Printf("Append only: %t\n", stats.Mode()&os.ModeAppend != 0)
	fmt.Printf("Device: %t\n", stats.Mode()&os.ModeDevice != 0)
	fmt.Printf("Symlink: %t\n", stats.Mode()&os.ModeSymlink != 0)
}
