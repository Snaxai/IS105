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

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("Kunne ikke åpne filen: text.txt %s\n", err)
		os.Exit(1)
	}

	size, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Panicf("failed at reading data from file: %s", err)
	}

	//Skriver ut stats til filen text.txt
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Kunne ikke finne stats: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Informasjon om filen", fileInfo.Name())
	fmt.Printf("Filstørrelse: %d bytes\n", len(size))
	i := len(size)
	f := float64(i)

	//Skriver ut størrelsen på text.txt filen som vi opprettet. Først blir størrelsen skrevet ut i bytes, deretter bruker vi denne
	//størrelsen til å skrive ut verdien i kibibytes, mibibytes og gibibytes.
	fmt.Printf("Filstørrelse: %f Kibibytes\n", f/1e3) //Tar utgangspunkt i 190 og flytter kommaet tre plasser til venstre vha. "f/1e3". For å skrive ut i kibibytes.
	fmt.Printf("Filstørrelse: %f Mibibytes\n", f/1e6) //Tar utgangspunkt i 190 og flytter kommaet seks plasser til venstre vha. "f/1e6". For å skrive ut i mibibytes.
	fmt.Printf("Filstørrelse: %f Gibigytes\n", f/1e9) //Tar utgangspunkt i 190 og flytter kommaet ni plasser til venstre vha. "f/1e9". For å skrive ut i gibibytes.

	fmt.Printf("Directory: %t\n", fileInfo.IsDir())
	fmt.Printf("Regular: %t\n", fileInfo.Mode().IsRegular())
	fmt.Printf("Unix permission bits: %s\n", fileInfo.Mode().String())
	fmt.Printf("Append only: %t\n", fileInfo.Mode()&os.ModeAppend != 0)
	fmt.Printf("Device file: %t\n", fileInfo.Mode()&os.ModeDevice != 0)
	fmt.Printf("Unix character device: %t\n", fileInfo.Mode()&os.ModeCharDevice != 0)
	fmt.Printf("Symbolic link: %t\n", fileInfo.Mode()&os.ModeSymlink != 0)
}
