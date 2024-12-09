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
	re2 := regexp.MustCompile("[0-9]+")
	res := 0
	operation := "mul"
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		s := re.FindAllStringSubmatch(line, -1)
		// fmt.Println(s)
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
	}
	fmt.Println("Result first star: ", res)
}
