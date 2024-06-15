package ascii

func JustNewLineHundler(Word *[]string) {
	EmptyElement := 0
	for _, w := range *Word {
		if w == "" {
			EmptyElement++
		}
	}
	if EmptyElement == len(*Word) {
		*Word = (*Word)[1:]
	}
}
