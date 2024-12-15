package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	r_wall       = '#'
	r_person     = '@'
	r_box        = 'O'
	r_empty      = '.'
	r_up         = '^'
	r_right      = '>'
	r_down       = 'v'
	r_left       = '<'
	s_directions = "^>v<"
)

type Person struct {
	Prow int
	Pcol int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	var input_map []string
	var movements []rune
	all_map_made := false
	var person Person
	// Split sequence into the map and movements
	for ri, line := range input {
		is_all_wall := true
		for ci, elem := range line {
			if strings.ContainsRune(s_directions, elem) {
				movements = append(movements, elem)
			}
			if !all_map_made {
				if elem == r_person {
					person = Person{Prow: ri, Pcol: ci}
				}
				if elem != '#' {
					is_all_wall = false
				}
			}
		}
		if !all_map_made {
			input_map = append(input_map, line)
		}
		if is_all_wall && ri > 0 {
			all_map_made = true
		}
	}
	fmt.Println(string(movements))

	for _, movement := range movements {
		cur_row := person.Prow
		cur_col := person.Pcol
		switch movement {
		case r_up:
			ri := cur_row
			for ; input_map[ri][cur_col] != r_wall && input_map[ri][cur_col] != r_empty; ri-- {
			}
			if input_map[ri][cur_col] == r_wall {
				continue
			}
			for person_switched := false; !person_switched; ri++ {
				cur_elem := input_map[ri][cur_col]
				next_elem := input_map[ri+1][cur_col]
				if next_elem == r_person {
					person_switched = true
					person.Prow = ri
				}
				input_map[ri] = replaceRuneInStringAtIndex(input_map[ri], cur_col, rune(next_elem))
				input_map[ri+1] = replaceRuneInStringAtIndex(input_map[ri+1], cur_col, rune(cur_elem))

			}
		case r_down:
			ri := cur_row
			for ; input_map[ri][cur_col] != r_wall && input_map[ri][cur_col] != r_empty; ri++ {
			}
			if input_map[ri][cur_col] == r_wall {
				continue
			}
			for person_switched := false; !person_switched; ri-- {
				cur_elem := input_map[ri][cur_col]
				next_elem := input_map[ri-1][cur_col]
				if next_elem == r_person {
					person_switched = true
					person.Prow = ri
				}
				input_map[ri] = replaceRuneInStringAtIndex(input_map[ri], cur_col, rune(next_elem))
				input_map[ri-1] = replaceRuneInStringAtIndex(input_map[ri-1], cur_col, rune(cur_elem))
			}
		case r_left:
			ci := cur_col
			for ; input_map[cur_row][ci] != r_wall && input_map[cur_row][ci] != r_empty; ci-- {
			}
			if input_map[cur_row][ci] == r_wall {
				continue
			}
			for person_switched := false; !person_switched; ci++ {
				cur_elem := input_map[cur_row][ci]
				next_elem := input_map[cur_row][ci+1]
				if next_elem == r_person {
					person_switched = true
					person.Pcol = ci
				}
				input_map[cur_row] = replaceRuneInStringAtIndex(input_map[cur_row], ci, rune(next_elem))
				input_map[cur_row] = replaceRuneInStringAtIndex(input_map[cur_row], ci+1, rune(cur_elem))
			}
		case r_right:
			ci := cur_col
			for ; input_map[cur_row][ci] != r_wall && input_map[cur_row][ci] != r_empty; ci++ {
			}
			if input_map[cur_row][ci] == r_wall {
				continue
			}
			for person_switched := false; !person_switched; ci-- {
				cur_elem := input_map[cur_row][ci]
				next_elem := input_map[cur_row][ci-1]
				if next_elem == r_person {
					person_switched = true
					person.Pcol = ci
				}
				input_map[cur_row] = replaceRuneInStringAtIndex(input_map[cur_row], ci, rune(next_elem))
				input_map[cur_row] = replaceRuneInStringAtIndex(input_map[cur_row], ci-1, rune(cur_elem))
			}
		}
	}

	total_gps := 0
	for ri, line := range input_map {
		for ci, elem := range line {
			if elem == r_box {
				total_gps += getGPSCoords(ri, ci)
			}
		}
	}

	fmt.Println("Total GPS: ", total_gps)
}

func getGPSCoords(row, col int) int {
	return (row * 100) + col
}

func printInputMap(input_map []string) {
	for _, line := range input_map {
		fmt.Println(line)
	}
}

func replaceRuneInStringAtIndex(str string, ind int, elem rune) string {
	ret_runes := []rune(str)
	ret_runes[ind] = elem
	return string(ret_runes)
}

func getInputAsLines() ([]string, error) {
	// Read in files
	// f, err := os.Open("test1.txt")
	// f, err := os.Open("test.txt")
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
