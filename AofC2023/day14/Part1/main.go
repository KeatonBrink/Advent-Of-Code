package main

import (
	"bufio"
	"fmt"
	"os"
)

var CYCLE_MEMORY = make(map[string][][]byte, 0)

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

func OneCycle(rock_map [][]byte) (ret_bytes [][]byte) {
	initial_byte_string := StringifySliceStrings(rock_map)

	val, ok := CYCLE_MEMORY[initial_byte_string]
	if ok {
		return val
	}

	CYCLE_MEMORY[initial_byte_string] = ret_bytes
	return
}

func StringifySliceStrings(bytes [][]byte) (str string) {
	for _, line := range bytes {
		str += string(line)
	}
	return
}

// Should not have done this
// func DeStringifyToBytes(str string, bytes_slice_len int) (bytes [][]byte) {
// 	for len(str) > 0 {
// 		var temp_bytes []byte
// 		for i := 0; i < bytes_slice_len; i++ {
// 			temp_bytes = append(temp_bytes, str[i])
// 		}
// 		bytes = append(bytes, temp_bytes)
// 		str = str[bytes_slice_len:]
// 	}
// 	return
// }
