package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

const (
	is_wall  = iota
	is_empty = iota
	is_mix   = iota
)

const (
	r_wall       = '#'
	r_person     = '@'
	r_box        = 'O'
	r_box_l      = '['
	r_box_r      = ']'
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

	var input_map [][]rune
	var movements []rune
	all_map_made := false
	var person Person
	// Split sequence into the map and movements
	for ri, line := range input {
		if !all_map_made {
			input_map = append(input_map, []rune{})
		}
		is_all_wall := true
		for ci, elem := range line {
			if strings.ContainsRune(s_directions, elem) {
				movements = append(movements, elem)
			}
			if !all_map_made {
				if elem == r_person {
					person = Person{Prow: ri, Pcol: ci * 2}
				}
				if elem != '#' {
					is_all_wall = false
				}
				if !all_map_made {
					switch elem {
					case r_wall:
						input_map[ri] = append(input_map[ri], r_wall)
						input_map[ri] = append(input_map[ri], r_wall)
					case r_empty:
						input_map[ri] = append(input_map[ri], r_empty)
						input_map[ri] = append(input_map[ri], r_empty)
					case r_person:
						input_map[ri] = append(input_map[ri], r_person)
						input_map[ri] = append(input_map[ri], r_empty)
					case r_box:
						input_map[ri] = append(input_map[ri], r_box_l)
						input_map[ri] = append(input_map[ri], r_box_r)
					}
				}
			}
		}
		if is_all_wall && ri > 0 {
			all_map_made = true
		}
	}
	fmt.Println(string(movements))

	printInputMap(input_map)

	for _, movement := range movements {
		objects_by_row := make(map[int][]int)
		cur_row := person.Prow
		cur_col := person.Pcol
		switch movement {
		case r_up:
			ri := cur_row - 1
			objects_by_row[cur_row] = []int{cur_col}
			impediments := is_mix
			for ; impediments == is_mix; ri-- {
				impediments = getMovementImpediments(input_map, objects_by_row, ri+1, ri)
			}
			if impediments == is_wall {
				continue
			}
			ri++
			for person_switched := false; !person_switched; ri++ {
				for _, ci := range objects_by_row[ri] {
					cur_elem := input_map[ri][ci]
					next_elem := input_map[ri-1][ci]
					if cur_elem == r_person {
						person_switched = true
						person.Prow--
					}
					input_map[ri][ci] = rune(next_elem)
					input_map[ri-1][ci] = rune(cur_elem)
				}

			}
		case r_down:
			ri := cur_row + 1
			objects_by_row[cur_row] = []int{cur_col}
			impediments := is_mix
			for ; impediments == is_mix; ri++ {
				impediments = getMovementImpediments(input_map, objects_by_row, ri-1, ri)
			}
			if impediments == is_wall {
				continue
			}
			ri--
			for person_switched := false; !person_switched; ri-- {
				for _, ci := range objects_by_row[ri] {
					cur_elem := input_map[ri][ci]
					next_elem := input_map[ri+1][ci]
					if cur_elem == r_person {
						person_switched = true
						person.Prow++
					}
					input_map[ri][ci] = rune(next_elem)
					input_map[ri+1][ci] = rune(cur_elem)
				}

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
				input_map[cur_row][ci] = rune(next_elem)
				input_map[cur_row][ci+1] = rune(cur_elem)
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
				input_map[cur_row][ci] = rune(next_elem)
				input_map[cur_row][ci-1] = rune(cur_elem)
			}
		}
		printInputMap(input_map)
		time.Sleep(30 * time.Millisecond)
	}

	total_gps := 0
	for ri, line := range input_map {
		for ci, elem := range line {
			if elem == r_box_l {
				total_gps += getGPSCoords(ri, ci)
			}
		}
	}

	fmt.Println("Total GPS: ", total_gps)
}

func getMovementImpediments(input_map [][]rune, objects_by_row map[int][]int, prev_ri, ri int) int {
	is_empty_space := true
	prev_obstruct_idxs := objects_by_row[prev_ri]
	for _, obj_ind := range prev_obstruct_idxs {
		elem := input_map[ri][obj_ind]
		switch elem {
		case r_wall:
			return is_wall
		case r_box_l:
			is_empty_space = false
			_, ok := objects_by_row[ri]
			if !ok {
				objects_by_row[ri] = []int{}
			}
			if !slices.Contains(objects_by_row[ri], obj_ind) {
				objects_by_row[ri] = append(objects_by_row[ri], obj_ind)
			}
			if !slices.Contains(objects_by_row[ri], obj_ind+1) {
				objects_by_row[ri] = append(objects_by_row[ri], obj_ind+1)
			}
		case r_box_r:
			is_empty_space = false
			_, ok := objects_by_row[ri]
			if !ok {
				objects_by_row[ri] = []int{}
			}
			if !slices.Contains(objects_by_row[ri], obj_ind) {
				objects_by_row[ri] = append(objects_by_row[ri], obj_ind)
			}
			if !slices.Contains(objects_by_row[ri], obj_ind-1) {
				objects_by_row[ri] = append(objects_by_row[ri], obj_ind-1)
			}
		}
	}
	if is_empty_space {
		return is_empty
	}
	return is_mix
}

func getGPSCoords(row, col int) int {
	return (row * 100) + col
}

func printInputMap(input_map [][]rune) {
	for _, line := range input_map {
		fmt.Println(string(line))
	}
}

func getInputAsLines() ([]string, error) {
	// Read in files
	// f, err := os.Open("test_simple.txt")
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
