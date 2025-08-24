package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "exampleinput_6.txt"
	// fmt.Print(len(text))
	// Loop over all lines in the file and print them.
	content, err := os.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println(lines[0])
	nchars := len(lines)
	var text [nchars][n]string
	for i := 0; i < len(lines); i++ {
		newArr := strings.Split(lines[i], "") // split on each char
		fmt.Println(newArr)
		fmt.Println(len(newArr))
		for j := 0; j < len(lines); j++ {
			text[i][j] = newArr[j]
		}
	}
	// fmt.Println(nUpdates, len(nUpdates))
}
