package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filename := "input_5.txt"
	// Open the file
	f, _ := os.Open(filename)
	// fmt.Print(len(text))
	scanner := bufio.NewScanner(f)
	var rules []string
	var left []int
	var right []int
	var update []string
	var nUpdates = map[int][]int{}
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else if strings.Contains(line, ",") {
			update = append(update, line)
		}
	}
	for _, rl := range rules {
		inter := strings.Split(rl, "|")
		le, _ := strconv.Atoi(inter[0])
		ri, _ := strconv.Atoi(inter[1])
		left = append(left, le)
		right = append(right, ri)
	}
	for i := 0; i < len(update); i++ {
		var numbers []int
		nums := strings.Split(update[i], ",")
		for j := 0; j < len(nums); j++ {
			n, _ := strconv.Atoi(nums[j])
			numbers = append(numbers, n)
		}
		nUpdates[i] = numbers
	}
	// fmt.Println(nUpdates, len(nUpdates))

	var RuleList = map[int][]int{}
	var checked []int
	getSorted(left, right, checked, RuleList)
	// fmt.Println("Does it work? ", RuleList)
	var notResults []int
	for t := 0; t < len(update); t++ {
		nums := strings.Split(update[t], ",")
		for i := len(nums) - 1; i > -1; i-- {
			n, _ := strconv.Atoi(nums[i])
			toCheck := RuleList[n]
			for j := i - 1; j > -1; j-- {
				nUp, _ := strconv.Atoi(nums[j])
				if slices.Contains(toCheck, nUp) {
					if slices.Contains(notResults, t) {

					} else {
						notResults = append(notResults, t)
					}
					break
				}
			}
		}
	}

	fmt.Println("Solution star 1: ", sumUp(update, notResults))
	fmt.Println("Solution star 2 stupid: ", sumUpStar2_2(update, notResults, RuleList))
}

func getSorted(left []int, right []int, checked []int, RuleList map[int][]int) {
	for i := 0; i < len(left); i++ {
		var nRule []int
		n := left[i]
		if slices.Contains(checked, n) {
		} else {
			nRule = append(nRule, right[i])
			for j := i + 1; j < len(left); j++ {
				if n == left[j] {
					nRule = append(nRule, right[j])
				}
			}
			RuleList[left[i]] = nRule
		}
		checked = append(checked, n)
	}
}
func sumUp(update []string, notResults []int) int {
	sum := 0
	for i := 0; i < len(update); i++ {
		if slices.Contains(notResults, i) {
			continue
		} else {
			nums := strings.Split(update[i], ",")
			mid := int(math.Ceil(float64(len(nums) / 2.0)))
			res, _ := strconv.Atoi(nums[mid])
			// fmt.Println(res)
			sum += res
		}
	}
	return sum
}

func sumUpStar2_2(update []string, notResults []int, RuleList map[int][]int) int {
	sum := 0
	for i := 0; i < len(update); i++ {
		var intNums []int
		if slices.Contains(notResults, i) {
			nums := strings.Split(update[i], ",")

			for j := 0; j < len(nums); j++ {
				hmm, _ := strconv.Atoi(nums[j])
				intNums = append(intNums, hmm)
			}

			// fmt.Println(intNums)
			// fmt.Println(score)
			loops := 13
			for kk := 0; kk < loops; kk++ {
				// correct := 0
				// fmt.Println("At the beginning of sorting ", intNums)
				for ii := len(intNums) - 1; ii > -1; ii-- {
					n := intNums[ii]
					toCheck := RuleList[n]
					// fmt.Println("Single element : ", n, " and a list of rules ", toCheck)

					for jj := ii - 1; jj > -1; jj-- {
						nUp := intNums[jj]
						if slices.Contains(toCheck, nUp) {
							// fmt.Println(nUp, "is in front of", n, "and a list of rules", toCheck)
							if kk == loops-1 {
								fmt.Println(nUp, "should be after", n)
							}
							intNums[jj] = n
							intNums[ii] = nUp
							// fmt.Println(correct, intNums)
							break
						}
					}
				}
				// fmt.Println("List at ", kk, "sorted ", intNums)
			}
			mid := int(math.Ceil(float64(len(intNums) / 2.0)))
			sum += intNums[mid]
		}
	}
	return sum
}
