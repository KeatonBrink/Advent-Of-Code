package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input_file_name := "input.txt"
	// input_file_name := "test_input.txt"

	read_file, err := os.Open(input_file_name)
	if err != nil {
		panic(err)
	}

	file_scanner := bufio.NewScanner(read_file)
	file_scanner.Split(bufio.ScanLines)

	var text []string

	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}

	seeds := strings.Split(text[0], " ")[1:]
	text = text[2:]

	first := true

	var transformers map[string]map[int]int

	var all_categories []string

	var current_category string

	var next_category string

	for _, line := range text {
		if unicode.IsLetter(rune(line[0])) {
			split_string := strings.Split(line, "-")
			current_category = split_string[0]
			next_category = split_string[2]
			transformers[current_category] = make(map[int]int)
			if first {
				all_categories = append(all_categories, current_category)
			} else {
				first = false
			}
			all_categories = append(all_categories, next_category)
		} else if unicode.IsDigit(rune(line[0])) {
			var source, destination, sd_range int
			found, err := fmt.Sscanf(line, "%d %d %d", &source, &destination, &sd_range)
			if err != nil {
				panic(err)
			}
			if found != 3 {
				fmt.Printf("Found only %d", found)
			}
		}
	}
}
