package unicode

import (
	"fmt"
)

const expression = "\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xc3\xb8\x72"
const is = "\xe2\x80\x9c\x6e\x6f\x72\xc3\xb0\x75\x72\x20\x6f\x67\x20\x73\x75\xc3\xb0\x75\x72\xe2\x80\x9d"
const jp = "\xe2\x80\x9c\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\xe2\x80\x9d"

// Gethexadesimal finner hexadesimal til en tekst
func Gethexadesimal(slice string) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%x ", slice[i])
	}
	fmt.Println(" ")
}

// Translate Kode for Oppgave 4a
func Translate(expression string, language string) string {
	if expression == "nord og sør" {
		if language == "jp" {
			fmt.Printf("%s \n", jp)
		}
		if language == "is" {
			fmt.Printf("%s \n", is)
		}
	}
	return ""
}

// UnicodeCodePointDemo Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// kan ikke tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
