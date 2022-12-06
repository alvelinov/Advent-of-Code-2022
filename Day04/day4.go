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

// for part 1
func completelyOverlappingIntervalsNum(intervals []string) int {
	re := regexp.MustCompile(`[\.]?[\d{2}]*`)

	var overlaps int = 0
	for _, str := range intervals {
		submatchall := re.FindAllString(str, -1)
		var arr = []int{0, 0, 0, 0}

		var i int = 0
		for _, element := range submatchall {
			n, _ := strconv.ParseFloat(element, 64)
			arr[i] = int(math.Abs(n))
			i++
		}

		if arr[0] >= arr[2] && arr[1] <= arr[3] || arr[2] >= arr[0] && arr[3] <= arr[1] {
			overlaps++
		}
	}

	return overlaps
}

// for part 2
func overlappingIntervalsNum(intervals []string) int {
	re := regexp.MustCompile(`[\.]?[\d{2}]*`)

	var overlaps int = 0
	for _, str := range intervals {
		submatchall := re.FindAllString(str, -1)
		var arr = []int{0, 0, 0, 0}

		var i int = 0
		for _, element := range submatchall {
			n, _ := strconv.ParseFloat(element, 64)
			arr[i] = int(math.Abs(n))
			i++
		}

		if arr[0] >= arr[2] && arr[0] <= arr[3] || arr[2] >= arr[0] && arr[2] <= arr[1] {
			overlaps++
		}
	}

	return overlaps
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

	intervals, err := readLines(inputFile)
	if err != nil {
		panic("Could not read input")
	}

	var result int
	if part == 1 {
		result = completelyOverlappingIntervalsNum(intervals)
	} else {
		result = overlappingIntervalsNum(intervals)
	}

	fmt.Println(result)
}
