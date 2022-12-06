package main

import (
	"os"
	"fmt"
	"bufio"
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

// returns your total score after all games have been played according to the strategy in part 1
func simulateGames1(moves []string) (int) {
	var sum int = 0
	var loss = 0
	var draw = 3
	var win = 6
	var rock = 1
	var paper = 2
	var scissors = 3
	for _, str := range moves {
		var left int
		switch str[0] {
		case 'A':
			left = rock
		case 'B':
			left = paper
		case 'C':
			left = scissors
		}

		var right int
		switch str[2] {
		case 'X':
			right = rock
			sum += rock
		case 'Y':
			right = paper
			sum += paper
		case 'Z':
			right = scissors
			sum += scissors
		}

		if left == right {
			sum += draw
		} else if left - right == -2 {
			sum += loss
		} else if left - right == 2 {
			sum += win
		} else if left - right == -1 {
			sum += win
		} else {
			sum += loss
		}
	}

	return sum
}

// returns your total score after all games have been played according to the strategy in part 2
func simulateGames2(moves []string) (int) {
	var sum int = 0
	var loss = 0
	var draw = 3
	var win = 6
	var rock = 1
	var paper = 2
	var scissors = 3
	for _, str := range moves {
		var left int
		switch str[0] {
		case 'A':
			left = rock
		case 'B':
			left = paper
		case 'C':
			left = scissors
		}

		switch str[2] {
		case 'X':
			if left == rock {
				sum += scissors
			} else if left == paper {
				sum += rock
			} else {
				sum += paper
			}
			sum += loss
		case 'Y':
			if left == rock {
				sum += rock
			} else if left == paper {
				sum += paper
			} else {
				sum += scissors
			}
			sum += draw
		case 'Z':
			if left == rock {
				sum += paper
			} else if left == paper {
				sum += scissors
			} else {
				sum += rock
			}
			sum += win
		}

	}

	return sum
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

	moves, err := readLines(inputFile)
	if err != nil {
		panic("Could not read input")
	}
	
	if part == 1 {
		var result = simulateGames1(moves)
		fmt.Println(result)
	} else {
		var result = simulateGames2(moves)
		fmt.Println(result)
	}
}