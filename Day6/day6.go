package main

import (
	"bufio"
	"fmt"
	"os"
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

func getStartOfPacketID(input string) (int) {
	var end int
	for i:=3; i<len(input); i++ {
		if input[i-3] != input[i-2] &&
		   		input[i-3] != input[i-1] &&
		   		input[i-3] != input[i]   &&
		   		input[i-2] != input[i-1] &&	
		   		input[i-2] != input[i]   &&
		   		input[i-1] != input[i] {
			end = i+1
			break
		}
	}

	return end
}

func getStartOfMessageID(input string) (int) {
	var end int
	for i:=13; i<len(input); i++ {
		var check [26]int
		var match bool = true
		for j:=i-13; j<=i; j++ {
			if check[input[j]-'a'] == 0 {
				check[input[j]-'a']++
			} else {
				match = false
				break
			}
		}
		if match {
			end = i+1
			break
		}
	}

	return end
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
	// In today's problem the input is a single line
	inputString := rawInput[0]

	if part == 1 {
		fmt.Println(getStartOfPacketID(inputString))
	} else {
		fmt.Println(getStartOfMessageID(inputString))
	}
	
}