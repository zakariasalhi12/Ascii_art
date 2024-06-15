package ascii

import (
	"fmt"
	"os"
)

const (
	ColorFlag   = "--color="
	JustifyFlag = "--align="
	OutputFlag  = "--output="
)

func Checker(words, ascii []string, option, flag, setlleters string, RGB RGB, trm_width int) {
	res := ""

	for i := 0; i < len(words); i++ {
		if flag == "" || flag == OutputFlag {
			res += Printer(words[i], ascii)
		} else if flag == JustifyFlag {
			Printer_Justify(words[i], ascii, flag, trm_width, option)
		} else if flag == ColorFlag {
			ColorPrinter(setlleters, words[i], ascii, RGB)
		}
	}
	if flag == "" || flag == ColorFlag && len(res) > 0 {
		fmt.Printf("%s", res)
	} else if flag == OutputFlag {
		if len(option) >= 4 && option[len(option)-4:] != ".txt" {
			fmt.Println(OutputMessageERR)
			return
		}
		file, _ := os.Create(option)
		file.WriteString(res)
		defer file.Close()
	}
}
