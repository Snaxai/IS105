package ascii

import (
	"fmt"
	"testing"

	"github.com/Snaxai/IS105/ICA02/Oppg1/ascii"
)

const PasserTestC = `"Hello :-)"`

func TestASCIIoppgC(t *testing.T) {
	for i := 0; i < len(PasserTestC); i++ {
		if PasserTestC[i] >= 126 { //fra og med 126 fordi det er utenfor normale ascii
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTestC[i])
		}
		fmt.Println(PasserTestC[i])
	}
}
func TestGreetingASCII(t *testing.T) {
	ascii.GreetingASCII()
}
