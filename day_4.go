package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "input_4.txt"
	// Open the file.
	f, _ := os.Open(filename)
	xmasLRcounter := 0
	xmasRLcounter := 0
	xmasUDcounter := 0
	xmasDUcounter := 0
	xmasRDLUcounter := 0
	xmasLDRUcounter := 0
	xmasLURDcounter := 0
	xmasRULDcounter := 0
	data, _ := os.ReadFile(filename)
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
	xmasDUcounter += readD2U(text, linelenght)
	xmasRDLUcounter += readDiaRDLU(text, linelenght)
	xmasLDRUcounter += readDiaLDRU(text, linelenght)
	xmasLURDcounter += readDiaLURD(text, linelenght)
	xmasRULDcounter += readDiaRULD(text, linelenght)

	fmt.Println("Left to right XMAS: ", xmasLRcounter)
	fmt.Println("Right to left XMAS: ", xmasRLcounter)
	fmt.Println("Up to down XMAS: ", xmasUDcounter)
	fmt.Println("Down to up XMAS: ", xmasDUcounter)
	fmt.Println("Dia R down to L up XMAS: ", xmasRDLUcounter)
	fmt.Println("Dia L down to R up XMAS: ", xmasLDRUcounter)
	fmt.Println("Dia L up to R down XMAS: ", xmasLURDcounter)
	fmt.Println("Dia R up to L down XMAS: ", xmasRULDcounter)
	fmt.Println("Finally we have : ", xmasDUcounter+xmasLRcounter+xmasRLcounter+xmasUDcounter+xmasRDLUcounter+xmasLDRUcounter+xmasLURDcounter+xmasRULDcounter)
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
		if string(text[i]) == "X" && i-3 > -1 {
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
						// fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
func readD2U(text string, linelength int) int {
	xmascount := 0
	nextline := linelength + 1
	for i := len(text) - 1; i > -1; i-- {
		if string(text[i]) == "X" && i-nextline*3 > -1 {
			if string(text[i-nextline]) == "M" {
				if string(text[i-2*nextline]) == "A" {
					if string(text[i-3*nextline]) == "S" {
						// fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
func readDiaRDLU(text string, linelength int) int {
	xmascount := 0
	nextline := linelength + 1
	for i := len(text) - 1; i > -1; i-- {
		if string(text[i]) == "X" && i-3-nextline*3 > -1 {
			if string(text[i-1-nextline]) == "M" {
				if string(text[i-2-2*nextline]) == "A" {
					if string(text[i-3-3*nextline]) == "S" {
						// fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
func readDiaLDRU(text string, linelength int) int {
	xmascount := 0
	nextline := linelength + 1
	for i := len(text) - 1; i > -1; i-- {
		if string(text[i]) == "X" && i+3-nextline*3 > -1 {
			if string(text[i+1-nextline]) == "M" {
				if string(text[i+2-2*nextline]) == "A" {
					if string(text[i+3-3*nextline]) == "S" {
						// fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
func readDiaLURD(text string, linelength int) int {
	xmascount := 0
	nextline := linelength + 1
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "X" && i+3+nextline*3 < len(text) {
			if string(text[i+1+nextline]) == "M" {
				if string(text[i+2+2*nextline]) == "A" {
					if string(text[i+3+3*nextline]) == "S" {
						// fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
func readDiaRULD(text string, linelength int) int {
	xmascount := 0
	nextline := linelength + 1
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "X" && i-3+nextline*3 < len(text) {
			if string(text[i-1+nextline]) == "M" {
				if string(text[i-2+2*nextline]) == "A" {
					if string(text[i-3+3*nextline]) == "S" {
						// fmt.Println("index", i, "Xmas here: ", string(text[i]), string(text[i+nextline]), string(text[i+2*nextline]), string(text[i+3*nextline]))
						xmascount += 1
					}
				}
			}
		}
	}
	return xmascount
}
