package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input_2.txt")

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	var nl []int
	var nl2 []int

	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var linlist []int
		// var revlist []int
		for i := 0; i < len(parts); i++ {
			num, _ := strconv.Atoi(parts[i])
			linlist = append(linlist, num)
		}
		sl := 0
		sl = saveUnsave(len(parts), linlist)
		nl = append(nl, saveUnsave(len(parts), linlist))
		if sl != 1 {

			for i := 0; i < len(parts); i++ {
				var slist []int
				dummy := slices.Clone(linlist)
				slist = slices.Delete(dummy, i, i+1)
				decide := saveUnsave(len(parts)-1, slist)
				if decide == 1 {
					nl2 = append(nl2, saveUnsave(len(parts)-1, slist))
					break
				}
			}
		}
	}
	n1 := sumWithLoop(nl)
	n2 := sumWithLoop(nl2)
	fmt.Println("First star: ", n1)
	fmt.Println("More than 0: ", n2)
	fmt.Println("Second star: ", n1+n2)
}
func sumWithLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func saveUnsave(n int, numbers []int) int {
	pcount := 0
	ncount := 0
	for i := 1; i < n; i++ {
		if (numbers[i-1]-numbers[i]) == 1 || (numbers[i-1]-numbers[i]) == 2 || (numbers[i-1]-numbers[i]) == 3 {
			pcount += 1
		} else if (numbers[i-1]-numbers[i]) == -1 || (numbers[i-1]-numbers[i]) == -2 || (numbers[i-1]-numbers[i]) == -3 {
			ncount += 1
		}
	}
	if n-1 == ncount || n-1 == pcount {
		return 1
	} else {
		return 0
	}

}
