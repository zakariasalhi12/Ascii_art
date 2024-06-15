package ascii

import (
	"strings"
)

var Files = map[string]string{
	"st": "standard",
	"sh": "shadow",
	"th": "thinkertoy",
}
//  check in the first arg if it flag
func Check_flag(s string) bool {
	if len(s) >= 8 && (s[:8] == JustifyFlag || s[:8] == ColorFlag) {
		return true
	} else if len(s) >= 9 && (s[:9] == OutputFlag) {
		return true
	}
	return false
}

// geting the flag the option 
func Get_flag(s string) (option, flag string) {
	if s[:8] == JustifyFlag {
		flag = JustifyFlag
		option = s[8:]
	} else if s[:8] == ColorFlag {
		flag = ColorFlag
		option = s[8:]
	}
	if len(s) > 8 && s[:9] == OutputFlag {
		flag = OutputFlag
		option = s[9:]
	}
	return flag, option
}
// ------------- check-in the banner if exsist return contentof it in arr are removet from args ---------------- 
func Check_banner(len_args int, File string, args []string) (int, string, []string) {
	for _, element := range Files {
		if strings.ToLower(args[len_args-1]) == element || strings.ToLower(args[len_args-1]) == element+".txt" {
			File = element + ".txt"
			args = args[:len_args-1]
			len_args--
			break
		}
	}

	return len_args, File, args
}
