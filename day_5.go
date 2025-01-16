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
	fmt.Println("Solution star 2 stupid: ", sumUpStar2(update, notResults, RuleList))
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
func sumUpStar2(update []string, notResults []int, RuleList map[int][]int) int {
	sum := 0
	for i := 0; i < len(update); i++ {
		var intNums []int
		if slices.Contains(notResults, i) {
			nums := strings.Split(update[i], ",")
			for j := 0; j < len(nums); j++ {
				hmm, _ := strconv.Atoi(nums[j])
				intNums = append(intNums, hmm)
			}
			score := 0
			for s := 1; s < len(intNums); s++ {
				score += len(intNums) - s
			}
			fmt.Println(score)
			for p := make([]int, len(intNums)); p[0] < len(p); nextPerm(p) {
				newList := getPerm(intNums, p)
				fmt.Println(newList)
				correct := 0
				for ii := len(newList) - 1; ii > -1; ii-- {
					n := newList[ii]
					toCheck := RuleList[n]
					// fmt.Println("Single element : ", n, " and a list of rules ", toCheck)
					for jj := ii - 1; jj > -1; jj-- {
						nUp := newList[jj]
						if slices.Contains(toCheck, nUp) {
							// fmt.Println(nUp, "is in front of", n, "and a list of rules", toCheck)
							// fmt.Println(nUp, "should be after", n)
							correct -= 10
							break
							// fmt.Println(correct, newList)
						} else {
							// fmt.Println("Is this sorted? ", newList)
							// fmt.Println(nUp, "should be in front of", n)
							correct += 1
							// fmt.Println(correct, nUp, n, newList)
						}
					}

					if correct == score {
						// fmt.Println("maybe?", newList)
						mid := int(math.Ceil(float64(len(newList) / 2.0)))
						sum += newList[mid]
						break
					}
				}
			}
		}
	}
	return sum
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
