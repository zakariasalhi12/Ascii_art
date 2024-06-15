package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/func"
)

var (
	RGB          ascii.RGB
	AsciiArtFile []string
	flag         string
	option       string
	trm_width    int
	Word         string
	Setletters   string
	File         = "standard.txt"
	args         = os.Args[1:]
	len_args     = len(args)
)

func main() {
	// ------------------no args case-----------------//
	if len_args < 1 {
		fmt.Println(ascii.JustifyMessageERR)
		return
	}
	// -----------------check-in the banner-------------------//
	len_args, File, args = ascii.Check_banner(len_args, File, args)

	// --------------------reding banner file--------------------------//
	AsciiArtFile := ascii.ReadFiles("Assci/" + File)
	// ----------getting the falg and removeing it from args---------------- //
	if ascii.Check_flag(args[0]) {
		flag, option = ascii.Get_flag(args[0])
		args = args[1:]
		len_args--
	}
	if err := ascii.Isvalid_args(len_args, flag); err != "" {
		fmt.Println(err)
		return
	}

	// --------------------------hundle color flag-------------------------- //

	if flag == ascii.ColorFlag {
		RGB.R, RGB.G, RGB.B = ascii.Get_RGB(option)
		if len_args > 1 {
			if ascii.AsciiCheker([]byte(args[0])) {
				return
			}
			if len(args[0]) == 1 {
				AsciiArtFile = ascii.Colorize(args[0], File, RGB)
				flag = ""
				option = ""
				args = args[1:]
				len_args--
			} else {
				Setletters = args[0]

				args = args[1:]
				len_args--
			}

		} else if len_args == 1 {
			AsciiArtFile = ascii.Colorize_ALL(File, RGB)
			flag = ""
			option = ""
		}

		// ------------------getting terminal size---------------------- //

	} else if flag == ascii.JustifyFlag {
		trm_width = ascii.GetTerminalSize()
		if len_args > 0 {
			args[0] = ascii.RemoveExtraSpaces(args[0])
		}

	}

	if flag == "" && len_args != 1 {
		fmt.Println(ascii.JustifyMessageERR)
		return
	}

	if len_args > 0 {
		Word = args[0]
	}

	Words := strings.Split(Word, "\\n")
	ascii.JustNewLineHundler(&Words)
	// reading File
	ascii.Checker(Words, AsciiArtFile, option, flag, Setletters , RGB, trm_width)
}
