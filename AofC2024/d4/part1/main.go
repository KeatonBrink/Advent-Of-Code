package main

import (
	"bufio"
	"fmt"
	"os"
)

var xmasLength int = 4

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	return_val := 0
	for i, line := range input {
		for j, letter := range line {
			if letter == 'X' {
				temp_val := xmasCountPosition(i, j, input)
				if temp_val > 0 {
					fmt.Println("i, j, count: ", i, j, temp_val)
				}
				return_val += temp_val
			}
		}
	}
	fmt.Println("Final Xmas Count: ", return_val)
}

func xmasCountPosition(i int, j int, input []string) int {
	ret_val := 0
	xmas := "XMAS"
	// Check up and left, then clockwise
	if i >= xmasLength-1 {
		if j >= xmasLength-1 {
			if input[i-1][j-1] == xmas[1] && input[i-2][j-2] == xmas[2] && input[i-3][j-3] == xmas[3] {
				ret_val += 1
			}
		}
		if input[i-1][j] == xmas[1] && input[i-2][j] == xmas[2] && input[i-3][j] == xmas[3] {
			ret_val += 1
		}
		if j <= len(input[i])-len(xmas) {
			if input[i-1][j+1] == xmas[1] && input[i-2][j+2] == xmas[2] && input[i-3][j+3] == xmas[3] {
				ret_val += 1
			}
		}
	}
	if j <= len(input[i])-len(xmas) {
		if input[i][j+1] == xmas[1] && input[i][j+2] == xmas[2] && input[i][j+3] == xmas[3] {
			ret_val += 1
		}
	}
	if i <= len(input)-len(xmas) {
		if j <= len(input[i])-len(xmas) {
			if input[i+1][j+1] == xmas[1] && input[i+2][j+2] == xmas[2] && input[i+3][j+3] == xmas[3] {
				ret_val += 1
			}
		}
		if input[i+1][j] == xmas[1] && input[i+2][j] == xmas[2] && input[i+3][j] == xmas[3] {
			ret_val += 1
		}
		if j >= xmasLength-1 {
			if input[i+1][j-1] == xmas[1] && input[i+2][j-2] == xmas[2] && input[i+3][j-3] == xmas[3] {
				ret_val += 1
			}
		}
	}
	if j >= xmasLength-1 {
		if input[i][j-1] == xmas[1] && input[i][j-2] == xmas[2] && input[i][j-3] == xmas[3] {
			ret_val += 1
		}
	}
	return ret_val
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
