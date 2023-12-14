package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	Row, Col int
}

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

	var rock_map [][]byte

	for _, line := range text {
		rock_map = append(rock_map, []byte(line))
	}

	for iterative_row := 1; iterative_row < len(rock_map); iterative_row++ {
		for column_of_row := 0; column_of_row < len(rock_map[iterative_row]); column_of_row++ {
			for temp_row := iterative_row; temp_row >= 1; temp_row-- {
				cur_val := &rock_map[temp_row][column_of_row]
				next_val := &rock_map[temp_row-1][column_of_row]
				if *cur_val == 'O' && *next_val == '.' {
					*cur_val = '.'
					*next_val = 'O'
				} else {
					break
				}
			}
		}
	}

	for _, line := range rock_map {
		println(string(line))
	}
	println()

	total_load := 0

	for iterative_row := 0; iterative_row < len(rock_map); iterative_row++ {
		for column_of_row := 0; column_of_row < len(rock_map[iterative_row]); column_of_row++ {
			cur_val := rock_map[iterative_row][column_of_row]
			if cur_val == 'O' {
				total_load += len(rock_map) - iterative_row
			}
		}
	}

	fmt.Printf("Score: %d\n", total_load)
}
