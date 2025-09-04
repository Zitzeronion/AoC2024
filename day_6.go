package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "exampleinput_6.txt"
	// fmt.Print(len(text))
	// Loop over all lines in the file and print them.
	content, err := os.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	n := len(lines) - 1
	// fmt.Println(lines[0])

	input := make([][]string, n)
	inputMod := make([][]string, n)
	path := make([][]int, n)

	for i := 0; i < n; i++ {
		newArr := strings.Split(lines[i], "") // split on each char

		// fmt.Println(newArr)
		// fmt.Println(len(newArr))
		// fmt.Println(newArr[2])
		for chars := 0; chars < len(newArr); chars++ {
			input[i] = append(input[i], newArr[chars])
			inputMod[i] = append(inputMod[i], newArr[chars])
			path[i] = append(path[i], 0)
		}

	}

	// fmt.Println(input[1])
	// fmt.Println(len(input[1]))
	guard := getPos(input)
	// Star 1
	updateGuard := moveGuard(guard, input, path)
	fmt.Println(updateGuard)
	// Star 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if inputMod[i][j] == "^" {
				inputMod[i][j] = "."
			}
		}
	}
	fmt.Println("Now the second star:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if inputMod[i][j] == "." {
				inputMod[i][j] = "#"
			}
			fmt.Println("Modification at: [", i, " ", j, "]")
			updateGuard2 := moveGuard(guard, inputMod, path)
			fmt.Println(updateGuard2)
		}
	}
	// fmt.Println(nUpdates, len(nUpdates))
}

func getPos(field [][]string) [2]int {
	n := len(field[1])
	pos := [2]int{0, 0}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if field[i][j] == "^" {
				pos[0] = i
				pos[1] = j
				field[i][j] = "."
			}
		}

	}
	return pos
}

func moveGuard(guard [2]int, field [][]string, path [][]int) [2]int {
	count := 0
	infloop := 0
	n := len(field[1])
	for {
		for i := guard[0] - 1; i > -1; i-- {
			if field[i][guard[1]] == "." {
				// Move horizontal
				guard[0] = i
				// fmt.Println(guard)
				path[i][guard[1]] = 1
			} else if field[i][guard[1]] == "#" {
				// fmt.Println("Up dead end", guard) // up dead end
				break
			} else if i == 0 {
				fmt.Println("Reached the border!")
				break
			}
			count++
		}
		if guard[0] == 0 || guard[0] == n-1 || guard[1] == 0 || guard[1] == n-1 {
			break
		}
		for i := guard[1] + 1; i < n; i++ {
			if field[guard[0]][i] == "." {
				// Move to the right
				guard[1] = i
				// fmt.Println(guard)
				path[guard[0]][i] = 1
			} else if field[guard[0]][i] == "#" {
				// fmt.Println("Right dead end", guard)
				break
			} else if guard[1] == n-1 {
				fmt.Println("Reached the border!")
				break
			}
			count++
		}
		if guard[0] == 0 || guard[0] == n-1 || guard[1] == 0 || guard[1] == n-1 {
			break
		}
		for i := guard[0] + 1; i < n; i++ {
			if field[i][guard[1]] == "." {
				// Move down
				guard[0] = i
				// fmt.Println(guard)
				path[i][guard[1]] = 1
			} else if field[i][guard[1]] == "#" {
				// fmt.Println("Bottom dead end", guard)
				break
			} else if guard[0] == n-1 {
				fmt.Println("Reached the border!")
				break
			}
			count++
		}
		if guard[0] == 0 || guard[0] == n-1 || guard[1] == 0 || guard[1] == n-1 {
			break
		}
		for i := guard[1] - 1; i > -1; i-- {
			if field[guard[0]][i] == "." {
				// Move to the left
				guard[1] = i
				// fmt.Println(guard)
				path[guard[0]][i] = 1
			} else if field[guard[0]][i] == "#" {
				// fmt.Println("Left dead end", guard)
				break
			} else if i == 0 {
				fmt.Println("Reached the border!")
				break
			}
			count++
		}
		if guard[0] == 0 || guard[0] == n-1 || guard[1] == 0 || guard[1] == n-1 {
			break
		}
		if count > 200 {
			fmt.Println(infloop)
			infloop++
			break
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += path[i][j]
		}
	}
	fmt.Println(sum, "and the count ", count)
	if infloop > 0 {
		fmt.Println("This infinite loop")
	}

	// fmt.Println(path)
	return guard
}
