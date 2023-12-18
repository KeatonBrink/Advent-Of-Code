package main

import (
	"bufio"
	"fmt"
	"os"
)

var CYCLE_MEMORY = make(map[string][][]byte, 0)

var ITERATION_MEMORY = make(map[string]int)

func main() {
	// input_file_name := "input.txt"
	input_file_name := "test_input.txt"

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

	for _, line := range rock_map {
		println(string(line))
	}
	println()

	for cur_cycle := 0; cur_cycle < 1000000000; cur_cycle++ {
		iteration_cycle := CheckDuplicate(rock_map, cur_cycle)
		if iteration_cycle == -1 {
			rock_map = OneCycle(rock_map, cur_cycle)
		} else {

		}
	}

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

func CheckDuplicate(rock_map [][]byte, iteration int) (duploop int) {
	initial_byte_string := StringifySliceStrings(rock_map)

	_, ok := CYCLE_MEMORY[initial_byte_string]
	if ok {
		return ITERATION_MEMORY[initial_byte_string]
	}

	return -1
}

func OneCycle(rock_map [][]byte, iteration int) (ret_bytes [][]byte) {
	initial_byte_string := StringifySliceStrings(rock_map)

	ret_bytes = make([][]byte, len(rock_map))
	for i := 0; i < len(rock_map); i++ {
		println("Making copy")
		ret_bytes[i] = make([]byte, len(rock_map[i]))
		copy(ret_bytes[i], rock_map[i])
	}

	for iterative_row := 1; iterative_row < len(ret_bytes); iterative_row++ {
		for column_of_row := 0; column_of_row < len(ret_bytes[iterative_row]); column_of_row++ {
			for temp_row := iterative_row; temp_row >= 1; temp_row-- {
				cur_val := &ret_bytes[temp_row][column_of_row]
				next_val := &ret_bytes[temp_row-1][column_of_row]
				if *cur_val == 'O' && *next_val == '.' {
					*cur_val = '.'
					*next_val = 'O'
				} else {
					break
				}
			}
		}
	}

	println(len(ret_bytes))
	for iterative_col := 1; iterative_col < len(ret_bytes[0]); iterative_col++ {
		for row_of_column := 0; row_of_column < len(ret_bytes); row_of_column++ {
			for temp_col := iterative_col; temp_col >= 1; temp_col-- {
				cur_val := &ret_bytes[row_of_column][temp_col]
				next_val := &ret_bytes[row_of_column][temp_col-1]
				if *cur_val == 'O' && *next_val == '.' {
					*cur_val = '.'
					*next_val = 'O'
				} else {
					break
				}
			}
		}
	}

	for iterative_row := len(ret_bytes) - 2; iterative_row > 0; iterative_row-- {
		for column_of_row := 0; column_of_row < len(ret_bytes[iterative_row]); column_of_row++ {
			for temp_row := iterative_row; temp_row < len(ret_bytes)-1; temp_row++ {
				cur_val := &ret_bytes[temp_row][column_of_row]
				next_val := &ret_bytes[temp_row+1][column_of_row]
				if *cur_val == 'O' && *next_val == '.' {
					*cur_val = '.'
					*next_val = 'O'
				} else {
					break
				}
			}
		}
	}

	for iterative_col := len(ret_bytes) - 2; iterative_col > 0; iterative_col-- {
		for row_of_column := 0; row_of_column < len(ret_bytes); row_of_column++ {
			for temp_col := iterative_col; temp_col < len(ret_bytes[0])-1; temp_col++ {
				cur_val := &ret_bytes[row_of_column][temp_col]
				next_val := &ret_bytes[row_of_column][temp_col+1]
				if *cur_val == 'O' && *next_val == '.' {
					*cur_val = '.'
					*next_val = 'O'
				} else {
					break
				}
			}
		}
	}

	CYCLE_MEMORY[initial_byte_string] = ret_bytes
	ITERATION_MEMORY[initial_byte_string] = iteration
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
