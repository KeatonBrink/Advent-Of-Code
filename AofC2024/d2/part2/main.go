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

		is_valid_line := false

		is_safe, problem_index := isGoodLine(numbers)
		if !is_safe {
			if problem_index < len(numbers)-1 {
				numbers = append(numbers[:problem_index], numbers[problem_index+1:]...)
			} else {
				numbers = numbers[:len(numbers)-1]
			}
			is_safe, _ = isGoodLine(numbers)
		}
		if is_safe {
			// fmt.Println("End Good Line: ", line, "\n\n\n")
			good_reports++
		} else {
			fmt.Println("End Bad line: ", line, "\n\n\n")
		}
	}

	fmt.Println("Good Reports: ", good_reports)
}

func isGoodLine(numbers []int) (bool, int) {
	increases := 0
	decreases := 0
	prev := 0
	is_first := true
	for i, elem := range numbers {
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
			temp := difference
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
				fmt.Println("Is Bad")
				fmt.Println("Difference: ", temp)
				fmt.Println("Increases ", increases, "\nDecreases", decreases)
				fmt.Println("Elements: ", elem, prev)
				fmt.Println("Index: ", i)
				fmt.Println("Line: ", numbers)
				if i == len(numbers)-1 {
					return false, i
				} else {
					return false, i - 1
				}
			}
		}
		prev = elem
	}
	return true, -1
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
