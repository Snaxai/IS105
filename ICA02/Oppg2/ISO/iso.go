package iso

import (
	"fmt"
)

// IterateOverExtendedASCIIStringLiteral For løkke
func IterateOverExtendedASCIIStringLiteral(slice string) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%X %q %b \n", slice[i], slice[i], slice[i])
	}
}
