package ascii

import (
	"bufio"
	"os"
)

// reading file and return array contatin every line of it file in seperated element
func ReadFiles(namefile string) []string {
	file, err := os.Open(namefile)
	if err != nil {
		panic("invalid file name ")
	}
	defer file.Close()
	read := bufio.NewScanner(file)
	var arr []string
	for read.Scan() {
		arr = append(arr, read.Text())
	}
	return arr
}
