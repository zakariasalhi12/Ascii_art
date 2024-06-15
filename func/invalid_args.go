package ascii

const (
	ColorMessageERR   = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something"
	JustifyMessageERR = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"
	FSMessageERR      = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
	OutputMessageERR  = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
)

func Isvalid_args(len_args int, flag string) string {
	if flag == OutputFlag && len_args != 1 {
		return OutputMessageERR
	} else if flag == ColorFlag && (len_args != 2 && len_args != 1) {
		return (ColorMessageERR)
	} else if flag == JustifyFlag && len_args != 1 {
		return (JustifyMessageERR)
	}
	return ""
}
