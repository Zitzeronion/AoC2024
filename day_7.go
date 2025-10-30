package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "exampleinput_7.txt"
	// fmt.Print(len(text))
	// Loop over all lines in the file and print them.
	content, err := os.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	n := len(lines) - 1
	fmt.Println(lines[0])

	for i := 0; i < n; i++ {
		leftside := strings.Split(lines[i], ":") // split on each char
		fmt.Println("Result", leftside[0])
		rightside := strings.Fields(leftside[1])
		fmt.Println("Terms", rightside)
		// fmt.Println(newArr[2])

	}

}
