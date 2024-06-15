package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	size := strings.Split(string(out), " ")
	termwidth, err := strconv.Atoi(size[1][:len(size[1])-1])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return termwidth
}
