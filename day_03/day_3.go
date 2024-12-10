package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Open the file.
	f, _ := os.Open("input_3.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// pat := regexp.MustCompile(`mul(.\([0-9]+,[0-9]+.)`)
	re := regexp.MustCompile(`[mul(]+\d[\d,]*[\d{2}]?[)]`)
	// re2 := regexp.MustCompile("[0-9]+")
	res := 0
	res2 := 0
	star := 2
	operation := "mul"
	linenumber := 0
	nextline := "whazup"
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		fmt.Println("We are reading line: ", linenumber)
		line := scanner.Text()
		s := re.FindAllStringSubmatch(line, -1)
		// fmt.Println(s)
		if star == 1 {
			res += checkString(s, operation)
		} else if star == 2 {
			star2 := strings.Split(line, "do")
			for i := 0; i < len(star2); i++ {
				star2[0] = nextline + star2[0]
				// fmt.Println(star2)
				if linenumber == 0 && i == 0 && star2[i][0:3] != "n't" {
					// fmt.Println("First line")
					// fmt.Println(star2[i])
					nextline = "()"
					s2 := re.FindAllStringSubmatch(star2[i], -1)
					res2 += checkString(s2, operation)
				} else if star2[i][0:2] == "()" {
					// fmt.Println("Multiply me")
					// fmt.Println(star2[i])
					nextline = "()"
					s2 := re.FindAllStringSubmatch(star2[i], -1)
					res2 += checkString(s2, operation)
				} else if star2[i][0:3] == "n't" {
					// fmt.Println("Throw me away")
					// fmt.Println(star2[i])
					nextline = "n't"
				} else {
					// fmt.Println("Think about")
				}
			}
		}
		linenumber++
	}
	if star == 1 {
		fmt.Println("Result first star: ", res)
	} else if star == 2 {
		fmt.Println("Result second star: ", res2)
	}
}

func checkString(s [][]string, operation string) int {
	res := 0
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i][0], operation) {
			parts := strings.Split(s[i][0], ",")
			if len(parts) == 2 {
				firstNum := re.FindString(parts[0])
				secNum := re.FindString(parts[1])
				num0, _ := strconv.Atoi(firstNum)
				num1, _ := strconv.Atoi(secNum)
				mul := num0 * num1
				res += mul
			} else {
				continue
			}
		}
	}
	// fmt.Println("Result : ", res)
	return res
}
