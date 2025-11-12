package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
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
	// fmt.Println(lines[0])

	for i := 0; i < n; i++ {
		leftside := strings.Split(lines[i], ":") // split on each char
		// fmt.Println("Result", leftside[0])
		result, _ := strconv.Atoi(leftside[0])
		fmt.Println(result)
		rightside := strings.Fields(leftside[1])
		var nums []int
		// fmt.Println("Terms", rightside)
		for j := 0; j < len(rightside); j++ {
			element, _ := strconv.Atoi(rightside[j])
			nums = append(nums, element)
		}
		sum := 0
		prod := 1
		for j := 0; j < len(rightside); j++ {
			sum += nums[j]
			prod *= nums[j]
		}
		if prod < result {
			continue
		}
		calc2 := 0
		for j := 0; j < len(rightside)-1; j++ {
			if j%2 == 0 {
				calc2 += nums[j] * nums[j+1]
			}
		}
		fmt.Println(sum, prod, calc2)
		fmt.Println(nums)

		// fmt.Println(newArr[2])

	}

}
