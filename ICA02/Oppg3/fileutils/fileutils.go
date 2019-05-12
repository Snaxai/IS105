package fileutils

import (
	"fmt"
	"io"
	"log"
	"os"
)

const tb = "\xEF\xDA\xA3\xD2\xD3\xCB\x0A\xEF\xDA\xA3\xD2\xD3\xCB\xC1\x0A\xEF" +
	"\x0A\x0A\xFE\xFD\x73\x6B\x61\x72\x0A\xFE\xFD\x73\x6B\x61\x72\x61\x6E\x61" +
	"\x0A\x0A\xF8\x79\x65\x73\x70\x65\x73\x69\x61\x6C\x69\x73\x74\x65\x6E\x0A"

//Print Konverterer og printer ut	const tb
func Print() {
	for i := 0; i < len(tb); i++ {
		//fmt.Printf("%q", tb[i])
		fmt.Printf("%c", tb[i])
	}
}

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
