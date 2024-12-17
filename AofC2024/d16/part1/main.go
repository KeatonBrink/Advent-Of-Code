package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	r_wall   = '#'
	r_space  = '.'
	r_person = 'S'
	r_end    = 'E'
)

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	for ri, line := range input {
		for ci, elem := range line {

		}
	}
}

func getInputAsLines() ([][]rune, error) {
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
	var text [][]rune
	for file_scanner.Scan() {
		text = append(text, []rune(file_scanner.Text()))
	}
	return text, nil
}
