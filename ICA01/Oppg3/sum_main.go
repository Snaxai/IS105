package main

import (
	"fmt"

	"github.com/Snaxai/IS105/ICA01/Oppg3/sum"
)

func main() {
	a := 2.5
	b := 3.5
	sumint8 := sum.SumInt8(55, 3)
	fmt.Println(sumint8)
	sumint32 := sum.SumInt32(5555, 2222)
	fmt.Println(sumint32)
	sumFloat64 := sum.SumFloat64(a, b)
	fmt.Println("Summen av", a, "og", b, "er", sumFloat64)
	sumuint32 := sum.SumuInt32(0, 700)
	fmt.Println(sumuint32)
	sumint64 := sum.SumInt64(-4222, 400)
	fmt.Println(sumint64)
}
