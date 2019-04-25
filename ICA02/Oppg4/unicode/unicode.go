package unicode

import (
	"fmt"
)

const nor = "\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xc3\xb8\x72"
const is = "\u0022\x6e\x6f\x72\u00f0\x75\x72\x20\x6f\x67\x20\x73\x75\u00f0\x75\x72\u0022"
const jp = "\u0022\xe5\x8c\x97\xE3\x81\xA8\xE5\x8D\x97\u0022"

// Gethexadesimal finner hexadesimal til en tekst
func Gethexadesimal(slice string) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%x \n", slice[i])
	}
}

// Translate Kode for Oppgave 4a
func Translate(expression string, language string) string {
	if expression == "nord og sør" {
		if language == "jp" {
			fmt.Printf("%s", jp)
		}
		if language == "is" {
			fmt.Printf("%s", is)
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
