package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	engraving  int
	next_stone *Stone
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	split_stones := strings.Fields(input[0])

	var stones []int
	// Create linked list stones
	for _, string_stone := range split_stones {
		stone_engraving, err := strconv.Atoi(string_stone)
		if err != nil {
			fmt.Println(err)
			return
		}
		stones = append(stones, stone_engraving)
	}

	// Map stone engraving to map of depths to stones
	var memo map[int]map[int][]int

	total_blinks := 25
	for blink := 0; blink < total_blinks; blink++ {
		// fmt.Println(stones)
		var new_stones []int
		for stone_index, stone := range stones {
			if stone == 0 {
				stones[stone_index] = 1
				continue
			}
			string_stone := strconv.Itoa(stone)
			if len(string_stone)%2 == 0 {
				a_string := string_stone[:len(string_stone)/2]
				b_string := string_stone[len(string_stone)/2:]
				a, err := strconv.Atoi(a_string)
				if err != nil {
					fmt.Println(err)
					return
				}
				b, err := strconv.Atoi(b_string)
				if err != nil {
					fmt.Println(err)
					return
				}
				stones[stone_index] = a
				new_stones = append(new_stones, b)
				continue
			}
			if len(string_stone)%2 == 1 {
				stones[stone_index] = stone * 2024
				continue
			}
		}
		stones = append(stones, new_stones...)
	}

	fmt.Println("Final: ", len(stones))
}

func descend(cur_blink, max_blink, self int, map_in_map map[int]map[int]int) int {
	if cur_blink == max_blink {
		return self
	}
	val, ok := map_in_map[self][cur_blink]
	if ok {
		return val
	}
	if self == 0 {
		temp_val := descend(cur_blink+1, max_blink, self+1, map_in_map)
		map_in_map[self][cur_blink] = temp_val
		return temp_val
	}
	string_stone := strconv.Itoa(self)
	if len(string_stone)%2 == 0 {
		a_string := string_stone[:len(string_stone)/2]
		b_string := string_stone[len(string_stone)/2:]
		a, err := strconv.Atoi(a_string)
		if err != nil {
			fmt.Println(err)
		}
		b, err := strconv.Atoi(b_string)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
	if len(string_stone)%2 == 1 {
		stones[stone_index] = stone * 2024
		continue
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
