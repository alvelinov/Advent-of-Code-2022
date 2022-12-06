package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
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


func sumCaloriesByElf(calories []string) ([]int, error) {
	caloriesByElf := []int{0}
	var lastID int = 0

	for _, str := range calories {
		if (str == "") {
			caloriesByElf = append(caloriesByElf, 0)
			lastID++
			continue
		}

		num, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}
		
		caloriesByElf[lastID] += num

	}

	return caloriesByElf, nil
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

	numStrings, err := readLines(inputFile)
	if err != nil {
		panic("Could not read input")
	}

	caloriesByElf, err := sumCaloriesByElf(numStrings)
	if err != nil {
		panic("Could not convert strings and sum numbers")
	}

	if part == 1 {
		var max int = 0
		for _, cal := range caloriesByElf {
			if cal > max {
				max = cal
			}
		}

		fmt.Println(max)
	} else {
		max := [3]int{0,0,0}
		for _, cal := range caloriesByElf {
			if cal > max[2] {
				max[0] = max[1]
				max[1] = max[2]
				max[2] = cal
			} else if cal > max[1] {
				max[0] = max[1]
				max[1] = cal
			} else if cal > max[0] {
				max[0] = cal
			}
		}

		fmt.Println(max[0] + max[1] + max[2])
	}

	return
}
