package ascii

import (
	"fmt"
	"strings"
)

func ColorPrinter(Setletters, word string, ascii []string, RGB RGB) {
	indexes := indexes(word, Setletters)
	colored := false
	counter := 0
	if len(word) == 0 {
		// protecting the case of empty element that means there new-line
		fmt.Println()
		return
	}
	byt := []byte(word)
	// Check if all words are ASCII.
	if AsciiCheker(byt) {
		return
	}
	// the looping for all the 8 lines
	for i := 0; i < 8; i++ {
		// this loop print the single line
		for j := 0; j < len(byt); j++ {
			// -----------------------------------------------//
			for k := 0; k < len(indexes); k++ {
				if j == indexes[k] {
					fmt.Print("\x1b[38;2;" + RGB.R + ";" + RGB.G + ";" + RGB.B + "m" + ascii[(int(byt[j])-32)*9+1+i] + "\x1b[0m")
					colored = true
					counter = 1
					break
				}
			}
			// -----------------------------------------------//

			if colored {
				colored = false
				continue
			}

			if counter != 0 && counter != len(Setletters) {
				fmt.Print("\x1b[38;2;" + RGB.R + ";" + RGB.G + ";" + RGB.B + "m" + ascii[(int(byt[j])-32)*9+1+i] + "\x1b[0m")
				counter++
				continue
			} else if counter == len(Setletters) {
				counter = 0
			}

			fmt.Print(ascii[(int(byt[j])-32)*9+1+i])
		}
		fmt.Println()
	}
}

func Printer(word string, ascii []string) string {
	data := ""

	if len(word) == 0 {
		// protecting the case of empty element that means there new-line
		return data + "\n"
	}
	byt := []byte(word)
	// Check if all words are ASCII.
	if AsciiCheker(byt) {
		return ""
	}
	// the looping for all the 8 lines
	for i := 0; i < 8; i++ {
		// this loop print the single line
		for j := 0; j < len(byt); j++ {
			data += ascii[(int(byt[j])-32)*9+1+i]
		}
		data += "\n"
	}
	return data
}

func Printer_Justify(word string, ascii []string, flag string, termsize int, option string) {
	WordLen, SpacesLen := WordAndSPacesLen(word, ascii)


	if len(word) == 0 {
		fmt.Println()
		return
	}
	byt := []byte(word)
	if AsciiCheker(byt) {
		return
	}

	for i := 0; i < 8; i++ {
		if option != "justify" {
			fmt.Print(AddsSpaces(termsize, WordLen, SpacesLen, option))
		}
		for j := 0; j < len(byt); j++ {
			if byt[j] == 32 && option == "justify" {
				fmt.Print(AddsSpaces(termsize, WordLen, SpacesLen, option))
			}
			fmt.Print(ascii[(int(byt[j])-32)*9+1+i])
		}
		fmt.Println()
	}
}

func indexes(word, Setletters string) []int {
	if len(Setletters) == 0 {
		return []int(nil)
	}
	index := []int{}
	offset := 0
	for {
		pos := strings.Index(word[offset:], Setletters)
		if pos == -1 {
			break
		}
		index = append(index, offset+pos)
		offset += pos + len(Setletters)
	}

	return index
}
