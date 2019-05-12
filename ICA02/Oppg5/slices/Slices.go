package slice

import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere
// Returnerer en slice av type []byte
func AllocateVar(b int) []byte {
	s := make([]byte, b, b)
	s[3] = 11
	s[5] = 69
	s = append(s, 12)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	return s
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
func AllocateMake(b int) []byte {
	d := make([]byte, b, b)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	return d
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	AllocateMake()
}

// CopySlice ???
