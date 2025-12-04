package main

import (
	"bufio"
	"fmt"
	"os"
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
	input = addBuffer(input, '.')
	roll_char := '@'
	total := 0
	for line_ind := 1; line_ind < len(input)-1; line_ind += 1 {
		line := input[line_ind]
		for char_ind := 1; char_ind < len(line)-1; char_ind += 1 {
			roll_count := 0
			char := line[char_ind]
			if char != byte(roll_char) {
				continue
			}
			roll_count += addIfRoll(input, line_ind-1, char_ind-1, roll_char)
			roll_count += addIfRoll(input, line_ind-1, char_ind, roll_char)
			roll_count += addIfRoll(input, line_ind-1, char_ind+1, roll_char)
			roll_count += addIfRoll(input, line_ind, char_ind-1, roll_char)
			roll_count += addIfRoll(input, line_ind, char_ind+1, roll_char)
			roll_count += addIfRoll(input, line_ind+1, char_ind-1, roll_char)
			roll_count += addIfRoll(input, line_ind+1, char_ind, roll_char)
			roll_count += addIfRoll(input, line_ind+1, char_ind+1, roll_char)

			if roll_count < 4 {
				// fmt.Println(line_ind, char_ind, input[line_ind][char_ind])
				// fmt.Println()

				total += 1
			}

		}
	}
	fmt.Println(total)
}

func addIfRoll(input []string, line_ind, char_ind int, roll_char rune) int {
	if input[line_ind][char_ind] == byte(roll_char) {
		// fmt.Println(line_ind, char_ind)
		return 1
	}
	return 0
}

func addBuffer(input []string, buffer_char rune) []string {
	output := make([]string, len(input))
	copy(output, input)
	for ind, line := range output {
		output[ind] = string(buffer_char) + line + string(buffer_char)
	}
	str_length := len(input[0]) + 2
	new_str := strings.Repeat(string(buffer_char), str_length)
	output = append(output, new_str)
	output = append([]string{new_str}, output...)

	return output
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
