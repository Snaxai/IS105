package fileutils

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FileToByteslice har en "string literal" som argument
// og returnerer en "slice"
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
	sizeOfSlice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, sizeOfSlice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	//Skriver ut symbolene i filen (unicode codepoints)
	fmt.Printf("%c\n", byteSlice)
	//Skriver ut symbolene i filen (Go syntax)
	fmt.Printf("%q\n", byteSlice)
	//Skriver ut desimalene til symbolene i filen
	fmt.Printf("", byteSlice)
	return byteSlice

}
