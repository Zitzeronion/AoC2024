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
	filename := "exampleinput_5.txt"
	// Open the file
	f, _ := os.Open(filename)
	// fmt.Print(len(text))
	scanner := bufio.NewScanner(f)
	var rules []string
	var left []int
	var right []int
	var update []string
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
					// fmt.Println("Why ", nUp, "this", toCheck, "at list ", t)
					break
				}
			}
		}
		// fmt.Println("UpateList ", t, "value ", easyNum)
	}
	fmt.Println("hmm: ", notResults)

	fmt.Println("Solution star 1: ", sumUp(update, notResults))
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
			// fmt.Println("Does it work? ", nRule, left[i])
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
func sumUpStar2(update []string, notResults []int) int {
	sum := 0
	for i := 0; i < len(update); i++ {
		if slices.Contains(notResults, i) {
			nums := strings.Split(update[i], ",")

			mid := int(math.Ceil(float64(len(nums) / 2.0)))
			res, _ := strconv.Atoi(nums[mid])
			// fmt.Println(res)
			sum += res
		} else {
			continue
		}
	}
	return sum
}
