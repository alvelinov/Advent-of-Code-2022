package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func parseInput(input []string) ([]Stack, [][]int) {
	var endId int = 0
	for id, str := range input {
		endId = id
		if str == "" {
			break
		}
	}

	stacksSlice := input[0:endId]
	instructions := [][]int{}
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	for i := endId + 1; i < len(input); i++ {
		submatchall := re.FindAllString(input[i], -1)
		var arr = []int{0, 0, 0}

		var j int = 0
		for _, element := range submatchall {
			n, _ := strconv.ParseFloat(element, 64)
			arr[j] = int(math.Abs(n))

			// Accounts for the way stacks are nubmered in the instructions
			// (I want them to start from 0 instead of 1)
			if j != 0 {
				arr[j]--
			}
			j++
		}
		instructions = append(instructions, arr)
	}

	submatchall := re.FindAllString(stacksSlice[len(stacksSlice)-1], -1)
	var stacksNum int
	for _, element := range submatchall {
		n, _ := strconv.ParseFloat(element, 64)
		stacksNum = int(math.Abs(n))
	}

	stacks := make([]Stack, stacksNum)
	for i := len(stacksSlice) - 2; i >= 0; i-- {
		for j := 1; j < len(stacksSlice[i]); j += 4 {
			if stacksSlice[i][j] >= 'A' && stacksSlice[i][j] <= 'Z' {
				stacks[j/4].Push(stacksSlice[i][j : j+1]) // pushes a single-letter string
			}
		}
	}

	return stacks, instructions
}

// Uses the rules from part 1
func followInstructions1(stacks []Stack, instructions [][]int) []Stack {
	for _, ins := range instructions {
		for i := ins[0]; i > 0; i-- {
			val, _ := stacks[ins[1]].Pop() // using '_' because the input is always valid
			stacks[ins[2]].Push(val)
		}
	}

	return stacks
}

// Uses the rules from part 2
func followInstructions2(stacks []Stack, instructions [][]int) []Stack {
	for _, ins := range instructions {
		cratesToMove := make([]string, ins[0])
		for i := ins[0]; i > 0; i-- {
			val, _ := stacks[ins[1]].Pop() // using '_' because the input is always valid
			cratesToMove[i-1] = val
		}
		stacks[ins[2]] = append(stacks[ins[2]], cratesToMove...)
	}

	return stacks
}

func main() {
	var part int
	fmt.Println("Which part of the challenge would you like to run? (1/2):")
	fmt.Scanln(&part)
	if part != 1 && part != 2 {
		panic("Invalid part")
	}

	var inputFile string
	fmt.Println("Write down the name of the input file you wish to fetch data from:")
	fmt.Scanln(&inputFile)

	rawInput, err := readLines(inputFile)
	if err != nil {
		panic("Could not read input")
	}

	stacks, instructions := parseInput(rawInput)

	if part == 1 {
		stacks = followInstructions1(stacks, instructions)
	} else {
		stacks = followInstructions2(stacks, instructions)
	}

	// Final output
	for _, stack := range stacks {
		val, _ := stack.Top()
		fmt.Print(val)
	}
}
