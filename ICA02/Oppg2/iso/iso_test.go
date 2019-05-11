package iso

import (
	"fmt"
	"testing"
)

func TestGreetingExtendedASCII(t *testing.T) {
	IterateOverExtendedASCIIStringLiteralB(GetExtendedASCIIStringLiteral())
}

const PasserTest = "Salut, ça va °-) Ça coute €50" +
	"Salut, ça va °-) Κοστίζει €50" + "Salut, ça va °-) Κοστίζει €50 Forstår du?"

func TestASCII(t *testing.T) {
	for i := 0; i < len(PasserTest); i++ {
		if PasserTest[i] >= 255 {
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTest[i])
		}
		if PasserTest[i] <= 0 {
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTest[i])
			fmt.Println(PasserTest[i])
		}
	}
}
