package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	Left, Right string
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

	var input_history [][]int
	for _, line := range text {
		temp_string_slice := strings.Split(line, " ")
		var temp_int_slice []int
		for _, el := range temp_string_slice {
			ret, err := strconv.Atoi(el)
			if err != nil {
				panic(err)
			}
			temp_int_slice = append(temp_int_slice, ret)
		}
		input_history = append(input_history, temp_int_slice)
	}

	// Parse line

	score := 0

	for _, line := range input_history {
		var individual_line_decomposition [][]int
		individual_line_decomposition = append(individual_line_decomposition, line)
		// println(line)
		for !AreAllZeroes(line) {
			var new_subhistory []int
			for i := 0; i < len(line)-1; i += 1 {
				new_subhistory = append(new_subhistory, line[i+1]-line[i])
			}
			individual_line_decomposition = append(individual_line_decomposition, new_subhistory)
			line = new_subhistory
		}
		var new_val int
		for depth := len(individual_line_decomposition) - 1; depth > 0; depth-- {
			higher_line := individual_line_decomposition[depth-1]
			lower_line := individual_line_decomposition[depth]
			new_val = higher_line[0] - lower_line[0]
			individual_line_decomposition[depth-1] = append([]int{new_val}, individual_line_decomposition[depth-1]...)
		}
		// println(new_val)
		score += new_val
	}

	fmt.Printf("Score found: %d\n", score)
}

func AreAllZeroes(line []int) bool {
	for _, element := range line {
		if element != 0 {
			return false
		}
	}
	return true
}
