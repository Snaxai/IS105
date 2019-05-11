package main

import (
	textscanner "github.com/Snaxai/IS105/ICA03/Oppg3/TextScanner"
	"github.com/Snaxai/IS105/ICA03/Oppg3/frequence"
)

func main() {
	frequence.FileToByteslice("pg100.txt")
	textscanner.FinnAlle("pg100.txt")
}
