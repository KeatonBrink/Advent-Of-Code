package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0

	for _, line := range input {
		largest, largest_ind := '0', -1
		for ind, char := range line {
			if largest < char {
				largest = char
				largest_ind = ind
			}
		}
		largest_double := "00"
		for ind, char := range line {
			if ind == largest_ind {
				continue
			}
			temp_str := "00"
			if ind < largest_ind {
				temp_str = string(char) + string(largest)
			} else {
				temp_str = string(largest) + string(char)
			}
			if largest_double < temp_str {
				largest_double = temp_str
			}
		}
		temp_val, err := strconv.Atoi(largest_double)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(temp_val)
		total += temp_val
	}
	fmt.Println(total)
}

func main_bad() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0
	// Find the two largest numbers
	for _, line := range input {
		first, second, first_ind, second_ind := '0', '0', -1, -1
		for ind, char := range line {
			if first < char {
				fmt.Println(string(first), first_ind, string(second), second_ind)
				second = first
				second_ind = first_ind
				first = char
				first_ind = ind
				fmt.Println(string(first), first_ind, string(second), second_ind)
			} else if second < char {
				fmt.Println(string(first), first_ind, string(second), second_ind)
				second = char
				second_ind = ind
				fmt.Println(string(first), first_ind, string(second), second_ind)
			}
		}
		var temp_str string
		if first_ind < second_ind {
			temp_str = string(first) + string(second)
		} else {
			temp_str = string(second) + string(first)
		}
		fmt.Println(temp_str)
		temp_val, err := strconv.Atoi(temp_str)
		if err != nil {
			fmt.Println(err)
			return
		}
		total += temp_val
	}
	fmt.Println(total)
}

func getInputAsLines(file_name string) ([]string, error) {
	// Read in files
	f, err := os.Open(file_name)
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
