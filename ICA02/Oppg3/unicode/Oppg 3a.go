package unicode

import (
	"fmt"
)

const unicode = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //x8c kan bli gjort om til xc6 og x98 kan gjort om til x7e
//for å skrie ut de riktige symbolene
const otherunicode = "\xc2\xbd\x3f\x3d\x3f\xe2\x8c\x98"

//SkrivUT skriver ut const unicode på %q
func SkrivUT() {
	for i := 0; i < len(unicode); i++ {
		fmt.Printf("%q ", unicode[i])
	}
}

//SkrivUTx skriver ut const unicode med %x
func SkrivUTx() {
	for i := 0; i < len(unicode); i++ {
		fmt.Printf("%x ", unicode[i])
	}
}

//Printunicode skriver ut const unicode
func Printunicode() {
	fmt.Println(unicode)
}

//SkrivUTotherunicode skriver ut const otherunicode
func SkrivUTotherunicode() {
	fmt.Println(otherunicode)
}
