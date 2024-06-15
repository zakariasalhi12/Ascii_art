package ascii

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RGB struct {
	R string
	G string
	B string
}

func Get_RGB(s string) (string, string, string) {
	s = change_color_name(s)
	check_valid_RGB(s)
	if s[0] == '#' {
		return hundle_hex(s[1:])
	}
	s = s[4 : len(s)-1]
	colors := strings.Split(s, ",")

	return colors[0], colors[1], colors[2]
}

func check_valid_RGB(s string) {
	if len(s) < 1 {
		os.Exit(0)
	}
	if s[0] == '#' && len(s) == 7 || len(s) == 4 {
		if !check_base(s[1:], "0123456789ABCDEFabcdef") {
			fmt.Println(ColorMessageERR)
			os.Exit(0)
		}
		return
	} else if len(s) > 3 && (s[:4] != "rgb(" || s[len(s)-1] != ')') {
		fmt.Println(ColorMessageERR)
		os.Exit(0)
	}
	if len(s) > 4 {
		s = s[4 : len(s)-1]
	}
	comma := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ',' && !(s[i] >= '0' && s[i] <= '9') {
			fmt.Println(ColorMessageERR)
			os.Exit(0)
		}
		if s[i] == ',' {
			comma++
		}
	}
	if comma != 2 {
		fmt.Println(ColorMessageERR)
		os.Exit(0)
	}
}

func hundle_hex(s string) (string, string, string) {
	if len(s) == 3 {
		x := ""
		for i := 0; i < 3; i++ {
			for j := 0; j < 2; j++ {
				x = x + string(s[i])
			}
		}
		s = x
	}
	var RGB RGB
	int_R, _ := strconv.ParseInt(string(s[0])+string(s[1]), 16, 64)
	int_G, _ := strconv.ParseInt(string(s[2])+string(s[3]), 16, 64)
	int_B, _ := strconv.ParseInt(string(s[4])+string(s[5]), 16, 64)

	RGB.R = strconv.Itoa(int(int_R))
	RGB.G = strconv.Itoa(int(int_G))
	RGB.B = strconv.Itoa(int(int_B))
	return RGB.R, RGB.G, RGB.B
}

func check_base(str, base string) bool {
	c := 0
	for i := 0; i < len(str); i++ {
		if is_in_base(str[i], base) {
			c++
		}
	}
	return c == len(str)
}

func is_in_base(number byte, base string) bool {
	for i := 0; i < len(base); i++ {
		if rune(base[i]) == rune(number) {
			return true
		}
	}
	return false
}

func change_color_name(s string) string {
	if value, ok := Colors[s]; ok {
		return value
	}
	return s
}
