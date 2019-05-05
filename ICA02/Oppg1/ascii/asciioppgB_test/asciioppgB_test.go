package ascii

import (
	"fmt"
	"testing"
)

//PasserTestA er en konstant som inneholder noen tegn
const PasserTestA = "\x00\x02\x07\x1d\x92\x96"

//TestASCIIA tester tegnene i konstanten PasserTestA
func TestASCIIA(t *testing.T) {
	for i := 0; i < len(PasserTestA); i++ {
		if PasserTestA[i] >= 126 {
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTestA[i])
		}
		fmt.Println(PasserTestA[i])
	}
}
