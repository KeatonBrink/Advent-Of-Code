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
	re := regexp.MustCompile("(mul\\(([1-9])([0-9]*),([1-9])([0-9]*)\\))|(do\\(\\))|(don't\\(\\))")
	output := 0
	is_active := true
	for _, line := range input {
		fmt.Println(line)
		instructions := re.FindAll([]byte(line), -1)
		fmt.Printf("%q\n", instructions)
		for _, instruction := range instructions {
			if string(instruction) == "do()" {
				is_active = true
			} else if string(instruction) == "don't()" {
				is_active = false
			} else if is_active {
				var a int
				var b int
				_, err := fmt.Sscanf(string(instruction), "mul(%d,%d)", &a, &b)
				if err != nil {
					fmt.Println(err)
					return
				}
				output += (a * b)
			}
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
