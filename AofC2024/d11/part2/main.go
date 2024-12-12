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
	memo := make(map[int]map[int]int)

	total_stones := 0
	total_blinks := 75
	for _, stone := range stones {
		total_stones += descend(0, total_blinks, stone, memo)
		// fmt.Println("z Total stones starting with x at y blinks", total_stones, stone, total_blinks)
	}

	fmt.Println("Final: ", total_stones)
}

func descend(cur_blink, max_blink, self int, map_in_map map[int]map[int]int) int {
	// fmt.Println(cur_blink, self)
	if cur_blink == max_blink {
		return 1
	}
	memo_val := getElemMapInMap(self, cur_blink, map_in_map)
	if memo_val >= 0 {
		return memo_val
	}
	if self == 0 {
		temp_val := descend(cur_blink+1, max_blink, self+1, map_in_map)
		_, ok := map_in_map[self]
		if !ok {
			map_in_map[self] = make(map[int]int)
		}
		map_in_map[self][cur_blink] = temp_val
		// fmt.Println(temp_val)
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
		a_ret := descend(cur_blink+1, max_blink, a, map_in_map)
		b, err := strconv.Atoi(b_string)
		if err != nil {
			fmt.Println(err)
		}
		b_ret := descend(cur_blink+1, max_blink, b, map_in_map)
		temp_val := a_ret + b_ret
		_, ok := map_in_map[self]
		if !ok {
			map_in_map[self] = make(map[int]int)
		}
		map_in_map[self][cur_blink] = temp_val
		// fmt.Println(temp_val)
		return temp_val
	} else {
		temp_val := descend(cur_blink+1, max_blink, self*2024, map_in_map)
		_, ok := map_in_map[self]
		if !ok {
			map_in_map[self] = make(map[int]int)
		}
		map_in_map[self][cur_blink] = temp_val
		// fmt.Println(temp_val)
		return temp_val
	}
}

func getElemMapInMap(engraving, blink int, map_in_map map[int]map[int]int) int {
	_, ok := map_in_map[engraving]
	if ok {
		val2, ok := map_in_map[engraving][blink]
		if ok {
			return val2
		}
	}
	return -1
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
