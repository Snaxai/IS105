package main

import (
	"github.com/Snaxai/IS105/ICA03/Oppg3/frequence"
)

func main() {
	frequence.FileinfoFrequence("pg100.txt")
	frequence.FileToByteslice("pg100.txt")
	frequence.Writetofile()
}
