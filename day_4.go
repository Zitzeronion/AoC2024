package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file.
	f, _ := os.Open("exampleinput_4.txt")
	xmasLRcounter := 0
	xmasRLcounter := 0
	xmasUDcounter := 0
	data, _ := os.ReadFile("exampleinput_4.txt")
	// Create a new Scanner for the file.
	text := string(data)

	// fmt.Print(len(text))
	scanner := bufio.NewScanner(f)
	linelenght := 0
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		linelenght = len(line)
		xmasLRcounter += readL2R(line)
		xmasRLcounter += readR2L(line)
	}
	xmasUDcounter += readU2D(text, linelenght)

	fmt.Println("Left to right XMAS: ", xmasLRcounter)
	fmt.Println("Right to left XMAS: ", xmasRLcounter)
	fmt.Println("Up to down XMAS: ", xmasUDcounter)
}
func readL2R(text string) int {
	xmascount := 0
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "X" && i+3 < len(text) {
			if string(text[i+1]) == "M" {
				if string(text[i+2]) == "A" {
					if string(text[i+3]) == "S" {
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}

func readR2L(text string) int {
	xmascount := 0
	for i := len(text) - 1; i > -1; i-- {
		if string(text[i]) == "X" && i-3 > 0 {
			if string(text[i-1]) == "M" {
				if string(text[i-2]) == "A" {
					if string(text[i-3]) == "S" {
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
func readU2D(text string, linelength int) int {
	xmascount := 0
	nextline := linelength + 1
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "X" && i+nextline*3 < len(text) {
			if string(text[i+nextline]) == "M" {
				if string(text[i+2*nextline]) == "A" {
					if string(text[i+3*nextline]) == "S" {
						fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
