package unicode

import (
	"fmt"
)

const unicode = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //x8c kan bli gjort om til xc6 og x98 kan gjort om til x7e
//for å skrie ut de riktige symbolene
func SkrivUT() {
	for i := 0; i < len(unicode); i++ {
		fmt.Printf("%q", unicode[i])
	}

}
