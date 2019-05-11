package main

import (
	textscanner "github.com/Snaxai/IS105/ICA03/Oppg3/TextScanner"
	"github.com/Snaxai/IS105/ICA03/Oppg3/frequence"
)

func main() {
	frequence.FileToByteslice("eks2.txt")
	textscanner.FinnAlle("eks2.txt")
}
