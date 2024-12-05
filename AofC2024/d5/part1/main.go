package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PageOrdering struct {
	previous *[]int
	future   *[]int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	is_blank := false

	orderings := make(map[int]PageOrdering)
	for _, line := range input {
		if strings.Contains(line, "|") {
			var first, second int
			_, err := fmt.Sscanf(line, "%d|%d", &first, &second)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, ok := orderings[first]
			if !ok {
				// orderings[first] = PageOrdering{previous: &[]int{second}}
				// pageOrdering := orderings[first]
				// *pageOrdering.future = append(*pageOrdering.future, second)
			}

		}
	}
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
