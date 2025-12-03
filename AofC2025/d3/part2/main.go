package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0

	// Take the previous answer, and
	for _, line := range input {
		bit_slice := make([]bool, len(line))

		for number_of_numbers_chosen := 0; number_of_numbers_chosen < 12; number_of_numbers_chosen += 1 {
			bit_slice = findNextVal(line, bit_slice)

		}
		temp_val, err := strconv.Atoi(genStrFromBoo(line, bit_slice, -1))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(temp_val)
		fmt.Println()
		total += temp_val
	}
	fmt.Println(total)
}

func findNextVal(line string, chosen_values []bool) []bool {
	best_slice := make([]bool, len(line))
	copy(best_slice, chosen_values)
	best_str := genStrFromBoo(line, chosen_values, -1)
	best_int := 0
	var err error
	if countTrueBits(chosen_values) > 0 {
		best_int, err = strconv.Atoi(best_str)
		if err != nil {
			fmt.Println(err)
		}
	}

	for ind, _ := range line {
		if !chosen_values[ind] {
			temp_str := genStrFromBoo(line, chosen_values, ind)
			temp_int, err := strconv.Atoi(temp_str)
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("comp strings", temp_str, best_str)
			if best_int < temp_int {
				// fmt.Println(best_str, temp_str)
				best_str = temp_str
				best_int, err = strconv.Atoi(best_str)
				if err != nil {
					fmt.Println(err)
				}
				copy(best_slice, chosen_values)
				best_slice[ind] = true
			}
		}
	}
	return best_slice
}

func genStrFromBoo(line string, chosen_values []bool, new_ind int) string {
	var ret_str string
	for ind, char := range line {
		if chosen_values[ind] == true || ind == new_ind {
			ret_str += string(char)
		}
	}
	return ret_str
}

func countTrueBits(bits []bool) int {
	total := 0
	for _, val := range bits {
		if val {
			total += 1
		}
	}
	return total
}

func getInputAsLines(file_name string) ([]string, error) {
	// Read in files
	f, err := os.Open(file_name)
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
