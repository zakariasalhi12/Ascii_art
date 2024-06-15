package ascii

import "os"

func WordAndSPacesLen(word string, ascci []string) (int, int) {
	counter := 0
	spaces := 0
	byt := []byte(word)
	if AsciiCheker(byt) {
		os.Exit(0)
	}
	// this loop print the single line
	for j := 0; j < len(byt); j++ {
		counter += len(ascci[(int(byt[j])-32)*9+1])
		if byt[j] == 32 {
			spaces++
		}
	}
	return counter, spaces
}
