package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	updates := []string{}

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
				orderings[first] = PageOrdering{previous: &[]int{}, future: &[]int{}}
			}
			_, ok = orderings[second]
			if !ok {
				orderings[second] = PageOrdering{previous: &[]int{}, future: &[]int{}}
			}
			*orderings[first].future = append(*orderings[first].future, second)
			*orderings[second].previous = append(*orderings[second].previous, first)
		}
		if strings.Contains(line, ",") {
			updates = append(updates, line)
		}
	}

	middle_page_sum := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")
		is_valid_ordering := true
		for cur_page_index, page_string := range pages {
			page, err := strconv.Atoi(page_string)
			if err != nil {
				fmt.Println(err)
				return
			}
			for prev_update_page_index := 0; prev_update_page_index < cur_page_index; prev_update_page_index++ {
				prev_update_page, err := strconv.Atoi(pages[prev_update_page_index])
				if err != nil {
					fmt.Println(err)
					return
				}
				if slices.Contains(*orderings[page].future, prev_update_page) {
					is_valid_ordering = false
					break
				}
			}
			for future_update_page_index := cur_page_index + 1; future_update_page_index < len(pages); future_update_page_index++ {
				future_update_page, err := strconv.Atoi(pages[future_update_page_index])
				if err != nil {
					fmt.Println(err)
					return
				}
				if slices.Contains(*orderings[page].previous, future_update_page) {
					is_valid_ordering = false
					break
				}
			}
			if !is_valid_ordering {
				break
			}
		}
		if is_valid_ordering {
			middle_page, err := strconv.Atoi(pages[len(pages)/2])
			if err != nil {
				fmt.Println(err)
				return
			}
			middle_page_sum += middle_page
		}
	}

	fmt.Println("Final: ", middle_page_sum)
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
