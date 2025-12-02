package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	input = strings.Split(input[0], ",")
	answer_set := make(map[int]bool)
	for _, val := range input {
		numbers := strings.Split(val, "-")
		start_r, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		end_r, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		for test_val := start_r; test_val <= end_r; test_val += 1 {
			test_val_string := strconv.Itoa(test_val)
			if len(test_val_string)%2 != 0 {
				continue
			}
			ind1 := 0
			for ind2 := 0; ind2 < len(test_val_string)/2; ind2 += 1 {
				sub_str := test_val_string[ind1 : ind2+1]
				fmt.Println(sub_str)
				if len(test_val_string)%len(sub_str) != 0 {
					continue
				}
				is_valid := true
				for mult := 1; mult*len(sub_str)+ind2 < len(test_val_string); mult += 1 {
					if mult > 1 {
						continue
					}
					test_sub := test_val_string[mult*len(sub_str) : ind2+mult*len(sub_str)+1]
					if sub_str != test_sub {
						is_valid = false
					}
				}
				if is_valid {
					fmt.Println("Is Valid", test_val, start_r, end_r)
					answer_set[test_val] = true
				}
			}
		}
	}
	total := 0
	for key, _ := range answer_set {
		total += key
	}

	fmt.Println(total)
}

func countDigitsString(num int) int {
	return len(strconv.Itoa(num))
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
