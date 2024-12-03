package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	// re := regexp.MustCompile("mul\\(([1-9]+)([0-9]?)([0-9]?),[:space:]([1-9]+)([0-9]?)([0-9]?)\\)")
	re := regexp.MustCompile("mul\\(([1-9])([0-9]*),([1-9])([0-9]*)\\)")
	output := 0
	for _, line := range input {
		fmt.Println(line)
		multipliers := re.FindAll([]byte(line), -1)
		fmt.Printf("%q\n", multipliers)
		for _, multi := range multipliers {
			var a int
			var b int
			_, err := fmt.Sscanf(string(multi), "mul(%d,%d)", &a, &b)
			if err != nil {
				fmt.Println(err)
				return
			}
			output += (a * b)
		}
	}
	fmt.Println("Final Output: ", output)
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
