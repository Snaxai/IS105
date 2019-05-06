package ascii

import (
	"fmt"
	"testing"

	"github.com/Snaxai/IS105/ICA02/Oppg1/ascii"
)

const PasserTest = `"Hello :-)"`

func TestASCII(t *testing.T) {
	for i := 0; i < len(PasserTest); i++ {
		if PasserTest[i] >= 126 { //til og med 126 fordi det er innenfor normale ascii
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTest[i])
		}
		fmt.Println(PasserTest[i])
	}
}

func TestGreetingASCII(t *testing.T) {
	ascii.GreetingASCII()
}
