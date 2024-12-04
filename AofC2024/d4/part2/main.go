package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	return_val := 0
	for i, line := range input {
		if i == 0 || i >= len(input)-1 {
			continue
		}
		for j, letter := range line {
			if j == 0 || j >= len(line)-1 {
				continue
			}
			if letter == 'A' {
				temp_val := isX_MAS(i, j, input)
				if temp_val {
					fmt.Println("Valid: i, j: ", i, j)
					return_val++
				}
			}
		}
	}
	fmt.Println("Final Xmas Count: ", return_val)
}

func isX_MAS(i int, j int, input []string) bool {
	ul := input[i-1][j-1]
	ur := input[i-1][j+1]
	dl := input[i+1][j-1]
	dr := input[i+1][j+1]
	if ul == 'M' && ur == 'M' && dl == 'S' && dr == 'S' {
		return true
	}
	if ul == 'M' && ur == 'S' && dl == 'M' && dr == 'S' {
		return true
	}
	if ul == 'S' && ur == 'M' && dl == 'S' && dr == 'M' {
		return true
	}
	if ul == 'S' && ur == 'S' && dl == 'M' && dr == 'M' {
		return true
	}
	return false
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
