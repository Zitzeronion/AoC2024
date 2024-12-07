package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file.
	f, _ := os.Open("input_1.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	var nl0 []int
	var nl1 []int
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		nums0, _ := strconv.Atoi(parts[0])
		nums1, _ := strconv.Atoi(parts[1])
		nl0 = append(nl0, nums0)
		nl1 = append(nl1, nums1)
		// fmt.Println(nums0, nums1)
	}
	sort.Ints(nl0)
	sort.Ints(nl1)
	var y []int
	var x []int
	for i := 0; i <= len(nl0)-1; i++ {
		count := 0
		x = append(x, Abs(nl0[i]-nl1[i]))
		for j := 0; j <= len(nl0)-1; j++ {
			if nl0[i] == nl1[j] {
				count += 1
			}
		}
		y = append(y, count)
	}
	fmt.Println("Answer first star")
	fmt.Println(sumWithLoop(x))
	// fmt.Println(y)
	muls := 0
	for i := 0; i <= len(nl0)-1; i++ {
		muls += nl0[i] * y[i]
	}
	fmt.Println("Answer second star")
	fmt.Println(muls)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sumWithLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
