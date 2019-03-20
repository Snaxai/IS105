package ascii

import (
	"fmt"
	"testing"
)

const PasserTest = `"Hello :-)"`

func TestASCII(t *testing.T) {
	for i := 0; i < len(PasserTest); i++ {
		if PasserTest[i] >= 126 {
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTest[i])
		}
		fmt.Println(PasserTest[i])
	}
}

func TestGreetingASCII(t *testing.T) {
	GreetingASCII()
}
