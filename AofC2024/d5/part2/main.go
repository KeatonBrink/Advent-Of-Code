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

	updates, orderings, err := getOrderings(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	middle_page_sum := 0
	for _, update := range updates {

		pages := strings.Split(update, ",")

		reordering_count, pages, err := getReorderingCount(pages, orderings)
		if err != nil {
			fmt.Println(err)
			return
		}
		if reordering_count > 0 {
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

func getReorderingCount(pages []string, orderings map[int]PageOrdering) (int, []string, error) {
	reordering_count := 0
	is_valid_ordering := false
	for !is_valid_ordering {
		var prev, cur, next int
		var err error
		is_valid_ordering, prev, cur, next, err = isValidOrdering(pages, orderings)
		if err != nil {
			fmt.Println(err)
			return -1, nil, err
		}
		if !is_valid_ordering {
			if prev > -1 {
				prev_page := pages[prev]
				pages[prev] = pages[cur]
				pages[cur] = prev_page
			} else if next > -1 {
				next_page := pages[next]
				pages[next] = pages[cur]
				pages[cur] = next_page
			} else {
				fmt.Println("Not good")

			}
			reordering_count++
		}
	}
	return reordering_count, pages, nil
}

func isValidOrdering(pages []string, orderings map[int]PageOrdering) (bool, int, int, int, error) {
	prev, cur, next := -1, -1, -1
	is_valid_ordering := true
	for cur_page_index, page_string := range pages {
		page, err := strconv.Atoi(page_string)
		if err != nil {
			fmt.Println(err)
			return false, -1, -1, -1, err
		}
		for prev_update_page_index := 0; prev_update_page_index < cur_page_index; prev_update_page_index++ {
			prev_update_page, err := strconv.Atoi(pages[prev_update_page_index])
			if err != nil {
				fmt.Println(err)
				return false, -1, -1, -1, err
			}
			if slices.Contains(*orderings[page].future, prev_update_page) {
				is_valid_ordering = false
				prev = prev_update_page_index
				cur = cur_page_index
				break
			}
		}
		for future_update_page_index := cur_page_index + 1; future_update_page_index < len(pages); future_update_page_index++ {
			future_update_page, err := strconv.Atoi(pages[future_update_page_index])
			if err != nil {
				fmt.Println(err)
				return false, -1, -1, -1, err
			}
			if slices.Contains(*orderings[page].previous, future_update_page) {
				is_valid_ordering = false
				next = future_update_page_index
				cur = cur_page_index
				break
			}
		}
		if !is_valid_ordering {
			break
		}
	}
	return is_valid_ordering, prev, cur, next, nil
}

func getOrderings(input []string) ([]string, map[int]PageOrdering, error) {
	updates := []string{}

	orderings := make(map[int]PageOrdering)
	for _, line := range input {
		if strings.Contains(line, "|") {
			var first, second int
			_, err := fmt.Sscanf(line, "%d|%d", &first, &second)
			if err != nil {
				fmt.Println(err)
				return nil, nil, err
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
	return updates, orderings, nil
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
