package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file.
	f, _ := os.Open("exampleinput_4.txt")
	xmascounter := 0
	// data, _ := os.ReadFile("exampleinput_4.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(string(line[0]))
		xmascounter += readL2R(line)
	}
	fmt.Println("Left to right XMAS: ", xmascounter)
}
func readL2R(text string) int {
	xmascount := 0
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "X" && i+3 < len(text) {
			if string(text[i+1]) == "M" && i+1 < len(text) {
				if string(text[i+2]) == "A" && i+2 < len(text) {
					if string(text[i+3]) == "S" && i+3 < len(text) {
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
