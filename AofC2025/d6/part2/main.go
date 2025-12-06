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

	var problem []int
	total := 0
	for col := len(input[0]) - 1; col >= 0; col -= 1 {
		var potential_num_str string
		for row := 0; row < len(input); row += 1 {
			fmt.Println(row, col, potential_num_str)
			if row == len(input)-1 {
				if potential_num_str != "" {
					potential_num, err := strconv.Atoi(potential_num_str)
					if err != nil {
						fmt.Println(err)
						return
					}
					problem = append(problem, potential_num)
				}
				if input[row][col] == '+' || input[row][col] == '*' {
					op := input[row][col]
					var temp_total int
					for ind2, num := range problem {
						if ind2 == 0 {
							temp_total = num
						} else {
							if op == '*' {
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
					problem = []int{}
				}
			} else if input[row][col] != ' ' {
				potential_num_str += string(input[row][col])
			}
		}
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
