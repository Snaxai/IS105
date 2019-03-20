package iso

import "fmt"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

// IterateOverASCIIStringLiteral tar en "string literal" som INN-data

func IteratOverExtendedASCIIStringLiteral() {
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X, %q, %b \n", i, i, i)
	}
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	return ""
}
