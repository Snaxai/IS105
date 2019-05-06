package ascii

import "fmt"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

//IterateOverASCIIStringLiteral skriver ut tegn
func IterateOverASCIIStringLiteral() {
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X %q %b\n", ascii[i], ascii[i], ascii[i])
	}
}

//GetASCIIStringLiteral returnerer ascii konstanten
func GetASCIIStringLiteral() string {
	return ascii
}

//GreetingASCII skriver ut en melding Hello :-)
func GreetingASCII() {
	tekst := `"Hello :-)"`
	for i := 0; i < len(tekst); i++ {
		fmt.Printf("%X %q %b\n", tekst[i], tekst[i], tekst[i])
	}
	fmt.Printf(tekst)
}
