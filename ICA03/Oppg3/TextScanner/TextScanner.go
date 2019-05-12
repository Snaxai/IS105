package textscanner

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//Items inneholder tegn som sjekkes
const Items = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,!?`

var biggest int
var andre int
var tredje int
var fjerde int
var femte int

var bokstav byte
var andreb byte
var tredjeb byte
var fjerdeb byte
var femteb byte

//FinnAlle lister ut de 5 mest brukte tegn som ligger i const Items
func FinnAlle() {
	valgtFil := flag.String("f", "", "filnavn")

	flag.Parse()
	if *valgtFil == "" {
		log.Fatal("Filen finnes ikke, bruk -f")
	}

	// Open file for reading
	file, err := os.Open(*valgtFil)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(Items); i++ {

		if bytes.Contains(byteSlice, []byte{Items[i]}) {
			v, _ := bytes.Count(byteSlice, []byte{Items[i]}), Items[i]
			if v > femte {
				femte = v
				femteb = Items[i]
				if fjerde < femte {
					fjerde = femte
					fjerdeb = femteb
					femte = 0

					if tredje < fjerde {
						tredje = fjerde
						tredjeb = fjerdeb
						fjerde = 0
						if andre < tredje {
							andre = tredje
							andreb = tredjeb
							tredje = 0
							if biggest < andre {
								biggest = andre
								bokstav = andreb
								andre = 0

							}

						}

					}

				}

			}

		}
	}
	fmt.Printf("Det er flest %q i denne filen og det er %d stk\n", bokstav, biggest)
	fmt.Printf("Det er nestflest %q i denne filen og det er %d stk\n", andreb, andre)
	fmt.Printf("Det er tredjflest %q i denne filen og det er %d stk\n", tredjeb, tredje)
	fmt.Printf("Det er fjerdflest %q i denne filen og det er %d stk\n", fjerdeb, fjerde)
	fmt.Printf("Det er femtflest %q i denne filen og det er %d stk\n", femteb, femte)
}

//Writetofilescanner asd
func Writetofilescanner() {
	tall1 := strconv.Itoa(biggest)
	tall2 := strconv.Itoa(andre)
	tall3 := strconv.Itoa(tredje)
	tall4 := strconv.Itoa(fjerde)
	tall5 := strconv.Itoa(femte)
	bokstav1 := string(bokstav)
	bokstav2 := string(andreb)
	bokstav3 := string(tredjeb)
	bokstav4 := string(fjerdeb)
	bokstav5 := string(femteb)
	err := ioutil.WriteFile("infoomfil.txt", []byte("Det er flest "+bokstav1+" i denne filen og det er "+tall1+" stk\n"+
		"Det er flest "+bokstav2+" i denne filen og det er "+tall2+" stk\n"+
		"Det er flest "+bokstav3+" i denne filen og det er "+tall3+" stk\n"+
		"Det er flest "+bokstav4+" i denne filen og det er "+tall4+" stk\n"+
		"Det er flest "+bokstav5+" i denne filen og det er "+tall5+" stk"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//FinnAlleTest brukes for å teste FinnAlle func
func FinnAlleTest(filename string) {
	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(Items); i++ {

		if bytes.Contains(byteSlice, []byte{Items[i]}) {
			v, _ := bytes.Count(byteSlice, []byte{Items[i]}), Items[i]
			if v > femte {
				femte = v
				femteb = Items[i]
				if fjerde < femte {
					fjerde = femte
					fjerdeb = femteb
					femte = 0

					if tredje < fjerde {
						tredje = fjerde
						tredjeb = fjerdeb
						fjerde = 0
						if andre < tredje {
							andre = tredje
							andreb = tredjeb
							tredje = 0
							if biggest < andre {
								biggest = andre
								bokstav = andreb
								andre = 0

							}

						}

					}

				}

			}

		}
	}
	fmt.Printf("Det er flest %q i denne filen og det er %d stk\n", bokstav, biggest)
	fmt.Printf("Det er nestflest %q i denne filen og det er %d stk\n", andreb, andre)
	fmt.Printf("Det er tredjflest %q i denne filen og det er %d stk\n", tredjeb, tredje)
	fmt.Printf("Det er fjerdflest %q i denne filen og det er %d stk\n", fjerdeb, fjerde)
	fmt.Printf("Det er femtflest %q i denne filen og det er %d stk\n", femteb, femte)
}
