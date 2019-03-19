package ascii

import (
	"fmt"
	"testing"
)

const PasserTest = "\x00\x02\x07\x1d\x92\x96"

func TestGreetingASCII(t *testing.T) {
	for i := 0; i < len(PasserTest); i++ {
		if PasserTest[i] >= 126 {
			t.Fail()
			fmt.Println("Value not a part of ascii in loop number", i, PasserTest[i])
		}
		fmt.Println(PasserTest[i])
	}
}
