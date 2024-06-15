package ascii

func AddsSpaces(termsize, wordlen, SpacesLen int, option string) string {
	spaces := 0
	res := ""
	if option == "center" {
		spaces = (termsize - wordlen) / 2
	}
	if option == "right" {
		spaces = (termsize - wordlen) - 1
	}
	if option == "left" {
		spaces = 0
	}
	if option == "justify" {
		if SpacesLen != 0 {
			spaces = (termsize - wordlen) / SpacesLen
		} else {
			spaces = 0
		}
	}
	for i := 0; i < spaces; i++ {
		res += " "
	}
	return res
}
