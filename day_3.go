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
	f, _ := os.Open("exampleinput_3.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// pat := regexp.MustCompile(`mul(.\([0-9]+,[0-9]+.)`)
	re := regexp.MustCompile(`[mul(]+\d[\d,]*[\d{2}]?[)]`)
	re2 := regexp.MustCompile("[0-9]+")
	res := 0
	star := 2
	operation := "mul"
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		star2 := strings.Split(line, "do")
		s := re.FindAllStringSubmatch(line, -1)
		// fmt.Println(s)
		if star == 1 {
			for i := 0; i < len(s); i++ {
				if strings.Contains(s[i][0], operation) {
					parts := strings.Split(s[i][0], ",")
					if len(parts) == 2 {
						firstNum := re2.FindString(parts[0])
						secNum := re2.FindString(parts[1])
						num0, _ := strconv.Atoi(firstNum)
						num1, _ := strconv.Atoi(secNum)
						mul := num0 * num1
						res += mul
					} else {
						continue
					}
				}
			}
		} else if star == 2 {
			for i := 0; i < len(star2); i++ {
				if star2[i][0:2] == "()" {
					fmt.Println("Multiply me")
					fmt.Println(star2[i])
				} else if star2[i][0:3] == "n't" {
					fmt.Println("Throw me away")
					fmt.Println(star2[i])
				} else {
					fmt.Println("Think about")
					fmt.Println(star2[i])
				}
				//fmt.Println(star2[i][0:3])
			}
		}
	}
	if star == 1 {
		fmt.Println("Result first star: ", res)
	} else if star == 2 {
		fmt.Println("Result second star: ", res)
	}
}
