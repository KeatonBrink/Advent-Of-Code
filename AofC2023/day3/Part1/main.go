package main

import (
	"bufio"
	"os"
	"unicode"
)

func main() {
	input_file_name := "input.txt"
	// input_file_name := "test_input.txt"


	read_file, err := os.Open(input_file_name)
	if err != nil {
		panic(err)
	}
	
	file_scanner := bufio.NewScanner(read_file)
	file_scanner.Split(bufio.ScanLines)

	var text_slice []string

	for file_scanner.Scan() {
		text_slice = append(text_slice, file_scanner.Text())
	}

	final_value := 0

	for line_ind, line := range(text_slice) {
		isNumValid := false
		curVal := -1
		for column, potential_num := range(line) {
			if unicode.IsDigit(potential_num) {
				if curVal !=  -1 {
					curVal = curVal * 10 + RuneToInt(potential_num)
				} else {
					curVal = RuneToInt(potential_num)
				}
				if IsSymbolAdjacent(text_slice, line_ind, column) {
					isNumValid = true
				}
			}
			if !unicode.IsDigit(potential_num) || column == len(line) -1 {
				if isNumValid {
					println("Valid number:", curVal)
					final_value += curVal
				}
				isNumValid = false
				curVal = -1
			}
		}
	}

	println("Final Sum: ", final_value)
}

func IsSymbolAdjacent(text []string, line_num, column int) bool {
	// Line above
	if line_num > 0 {
		// left up
		if column > 0 && IsSymbol(rune(text[line_num-1][column-1])) {
			return true
		}
		// middle up
		if IsSymbol(rune(text[line_num-1][column])) {
			return true
		}
		// right up
		if column < len(text[line_num - 1]) - 1 && IsSymbol(rune(text[line_num-1][column+1])) {
			return true
		}
	}
	// line below
	if line_num < len(text) - 1{
		// left below
		if column > 0 && IsSymbol(rune(text[line_num+1][column-1])) {
			return true
		}
		// middle below
		if IsSymbol(rune(text[line_num+1][column])) {
			return true
		}
		// right below
		if column < len(text[line_num + 1]) - 1 && IsSymbol(rune(text[line_num+1][column+1])) {
			return true
		}
	}
	// left
	if column > 0 && IsSymbol(rune(text[line_num][column-1])) {
		return true
	}
	// right
	if column < len(text[line_num]) - 1 && IsSymbol(rune(text[line_num][column + 1])) {
		return true
	}
	return false
}

func IsSymbol(character rune) bool {
	return !unicode.IsDigit(character) && character != '.'
}

func RuneToInt(character rune) int {
	return int(character - '0')
}