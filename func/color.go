package ascii

func Colorize(colorized, File_name string, RGB RGB) []string {
	arr := ReadFiles("Assci/" + File_name)
	chars := []rune(colorized)
	byt := []byte(colorized)
	for i := 0; i < len(chars); i++ {
		for j := 0; j < 8; j++ {
			x := (int(byt[i])-32)*9 + 1 + j
			arr[x] = "\x1b[38;2;" + RGB.R + ";" + RGB.G + ";" + RGB.B + "m" + arr[x] + "\x1b[0m"
		}
	}
	return arr
}

func Colorize_ALL(File_name string, RGB RGB) []string {
	arr := ReadFiles("Assci/" + File_name)
	for i := 0; i < len(arr); i++ {
		arr[i] = "\x1b[38;2;" + RGB.R + ";" + RGB.G + ";" + RGB.B + "m" + arr[i] + "\x1b[0m"
	}
	return arr
}
