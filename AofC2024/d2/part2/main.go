package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	good_reports := 0
	tests := 0

	for _, line := range input {
		string_ints := strings.Split(line, " ")
		var numbers []int
		for _, single_string := range string_ints {
			single_int, err := strconv.Atoi(single_string)
			if err != nil {
				fmt.Println(err)
				return
			}
			numbers = append(numbers, single_int)
		}

		is_valid_line := isGoodLine(numbers)
		tests++

		if !is_valid_line {
			for i := 0; i < len(numbers); i++ {
				copy_numbers := copyLineAndRemoveAtIndex(numbers, i)
				is_valid_line = isGoodLine(copy_numbers)
				tests++
				if is_valid_line {
					break
				}
			}
		}
		if is_valid_line {
			good_reports++
		}
	}

	fmt.Println("Good Reports: ", good_reports, " Tests: ", tests)
}

func copyLineAndRemoveAtIndex(numbers []int, index int) []int {
	copy_numbers := make([]int, len(numbers))
	copy(copy_numbers, numbers)
	if index == 0 {
		copy_numbers = copy_numbers[1:]
	} else if index == len(numbers)-1 {
		copy_numbers = copy_numbers[:len(copy_numbers)-1]
	} else {
		copy_numbers = append(copy_numbers[:index], numbers[index+1:]...)
	}
	fmt.Println("Copied: ", copy_numbers, " Original: ", numbers)
	return copy_numbers
}

func isGoodLine(numbers []int) bool {
	increases := 0
	decreases := 0
	prev := 0
	is_first := true
	for _, elem := range numbers {
		is_safe := true
		if is_first {
			is_first = false
		} else {
			difference := elem - prev
			// Check if increasing or decreasing
			if difference > 0 {
				increases++
			} else if difference < 0 {
				decreases++
			}
			// temp := difference
			difference = int(math.Abs(float64(difference)))
			// difference is less than 1 or greater than three
			// Which is bad
			if difference < 1 || difference > 3 {
				is_safe = false
			}
			if increases > 0 && decreases > 0 {
				is_safe = false
			}
			if !is_safe {
				// fmt.Println("Is Bad")
				// fmt.Println("Difference: ", temp)
				// fmt.Println("Increases ", increases, "\nDecreases", decreases)
				// fmt.Println("Elements: ", elem, prev)
				// fmt.Println("Index: ", i)
				// fmt.Println("Line: ", numbers)
				return false
			}
		}
		prev = elem
	}
	return true
}

func getInputAsLines() ([]string, error) {
	// Read in files
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Create scanner
	file_scanner := bufio.NewScanner(f)
	file_scanner.Split(bufio.ScanLines)

	// Get lines
	var text []string
	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}
	return text, nil
}
