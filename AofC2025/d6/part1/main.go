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
	var problems [][]int
	var operations []string
	for ind1, line := range input {
		line_fields := strings.Fields(line)
		if line_fields[0] == "+" || line_fields[0] == "*" {
			operations = line_fields
			break
		}
		for ind2, field := range line_fields {
			cur_int, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				return
			}
			if ind1 == 0 {
				problems = append(problems, []int{cur_int})
			} else {
				// fmt.Println(problems)
				// fmt.Println(ind2, cur_int)
				problems[ind2] = append(problems[ind2], cur_int)
			}
		}
	}
	total := 0
	for ind1, op := range operations {
		nums := problems[ind1]

		var temp_total int
		for ind2, num := range nums {
			if ind2 == 0 {
				temp_total = num
			} else {
				if op == "*" {
					temp_total *= num
					if temp_total == 0 {
						break
					}
				} else {
					temp_total += num
				}
			}
		}
		total += temp_total
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
