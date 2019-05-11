package frequence

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//FileinfoFrequence skriver ut info om fil
func FileinfoFrequence() {

	valgtFil := flag.String("f", "", "filnavn")

	flag.Parse()
	if *valgtFil == "" {
		log.Fatal("Filen finnes ikke, bruk -f")
	}
	file, err := os.Open(*valgtFil)
	if err != nil {
		fmt.Printf("Kunne ikke åpne filen: "+*valgtFil+" %s\n", err)
		os.Exit(1)
	}
	size, err := ioutil.ReadFile(*valgtFil)
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

//FileToByteslice finner linjeskift
func FileToByteslice(filename string) []byte {
	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	byteSlice := make([]byte, size_of_slice)

	lineBreak2 := ([]byte("\x0A"))
	lineBreak1 := ([]byte("\x0D\x0A"))

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	if bytes.Contains(byteSlice, lineBreak1) {
		fmt.Println("Det er", bytes.Count(byteSlice, lineBreak1), "linjeskift i denne filen.(Carriage return og lineshift)")
	} else if bytes.Contains(byteSlice, lineBreak2) {
		fmt.Println("Det er", bytes.Count(byteSlice, lineBreak2), "linjeskift i denne filen.(Lineshift)")
	} else {
		fmt.Println("Det er ingen linjeskift i filen.")
	}
	return byteSlice
}

func Writetofile() {
	err := ioutil.WriteFile("info.txt", []byte("asdasd\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
