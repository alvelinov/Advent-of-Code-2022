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


func getCommonItem1(str1 string, str2 string) (rune){
	for _, char1 := range str1 {
		for _, char2 := range str2 {
			if char1 == char2 {
				return char1
			}
		}
	}

	// the input is correct and this line will never execute
	return ' '
}

func makeLetterTracker(str string) ([52]bool) {
	var arr [52]bool
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			arr[char - 'a'] = true
		} else {
			arr[char - 'A' + 26] = true
		}
	}

	return arr
}

func getCommonItem2(str1 string, str2 string, str3 string) (rune) {
	var arr1 = makeLetterTracker(str1)
	var arr2 = makeLetterTracker(str2)
	var arr3 = makeLetterTracker(str3)
	
	var item rune
	for i:=0; i<52; i++ {
		if arr1[i] && arr1[i] == arr2[i] && arr2[i] == arr3[i] {
			if i > 25 {
				item = 'A' + int32(i) - 26
			} else {
				item = 'a' + int32(i)
			}
		}
	}

	return item
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

	rucksacks, err := readLines(inputFile)
	if err != nil {
		panic("Could not read input")
	}

	var prioritiesSum int = 0
	if part == 1 {
		for _, rucksack := range rucksacks {
			s1 := rucksack[0:len(rucksack)/2]
			s2 := rucksack[len(rucksack)/2:]
			var item = getCommonItem1(s1, s2)
		
			if item >= 'a' && item <= 'z' {
				prioritiesSum += int(item) - int('a') + 1
			} else {
				prioritiesSum += int(item) - int('A') + 27
			}
		}
	} else {
		for i:=0; i<len(rucksacks); i+=3 {
			var item = getCommonItem2(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
			
			if item >= 'a' && item <= 'z' {
				prioritiesSum += int(item) - int('a') + 1
			} else {
				prioritiesSum += int(item) - int('A') + 27
			}
		}
	}

	fmt.Println(prioritiesSum)

}