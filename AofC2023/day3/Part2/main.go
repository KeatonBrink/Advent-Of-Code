package main

import (
	"bufio"
	"os"
	"strconv"
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

	var text []string

	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}

	final_value := 0

	for line_num, line := range(text) {
		for column, potential_num := range(line) {
			// Search for *
			if potential_num == '*' {
				// 1 and 2 gear
				gear1 := ""
				gear2 := ""
				// Check all surrounding spots
				// Line above
				if line_num > 0 {
					// left up
					if column > 0 && unicode.IsDigit(rune(text[line_num-1][column-1])) {
						gear1, gear2 = UpdateGears(gear1, gear2, line_num-1, column - 1, text)
					}
					// middle up
					if unicode.IsDigit(rune(text[line_num-1][column])) {
						gear1, gear2 = UpdateGears(gear1, gear2, line_num-1, column, text)

					}
					// right up
					if column < len(text[line_num - 1]) - 1 && unicode.IsDigit(rune(text[line_num-1][column+1])) {
						gear1, gear2 = UpdateGears(gear1, gear2, line_num-1, column + 1, text)
					}
				}
				// line below
				if line_num < len(text) - 1{
					// left below
					if column > 0 && unicode.IsDigit(rune(text[line_num+1][column-1])) {
						gear1, gear2 = UpdateGears(gear1, gear2, line_num + 1, column - 1, text)
					}
					// middle below
					if unicode.IsDigit(rune(text[line_num+1][column])) {
						gear1, gear2 = UpdateGears(gear1, gear2, line_num + 1, column, text)

					}
					// right below
					if column < len(text[line_num + 1]) - 1 && unicode.IsDigit(rune(text[line_num+1][column+1])) {
						gear1, gear2 = UpdateGears(gear1, gear2, line_num + 1, column + 1, text)
					}
				}
				// left
				if column > 0 && unicode.IsDigit(rune(text[line_num][column-1])) {
					gear1, gear2 = UpdateGears(gear1, gear2, line_num, column - 1, text)
				}
				// right
				if column < len(text[line_num]) - 1 && unicode.IsDigit(rune(text[line_num][column + 1])) {
					gear1, gear2 = UpdateGears(gear1, gear2, line_num, column + 1, text)
				}
				if gear1 != "" && gear2 != "" {
					gear1_int, err := strconv.Atoi(gear1)
					if err != nil {
						panic(err)
					}
					gear2_int, err := strconv.Atoi(gear2)
					if err != nil {
						panic(err)
					}
					final_value += gear1_int * gear2_int
				}
			}
		}
	}

	println("Final Sum: ", final_value)
}

func UpdateGears(gear1, gear2 string, row, col int, text []string) (string, string) {
	temp_gear := string(text[row][col])
	temp_col := col
	for temp_col > 0 && unicode.IsDigit(rune(text[row][temp_col-1])) {
		temp_col -= 1
		temp_gear = string(text[row][temp_col]) + temp_gear
	}
	temp_col = col
	for temp_col < len(text[row]) - 1 && unicode.IsDigit(rune(text[row][temp_col+1])) {
		temp_col += 1
		temp_gear += string(text[row][temp_col])
	}
	if gear1 == "" {
		return temp_gear, gear2
	}
	if temp_gear != gear1 && temp_gear != gear2 {
		return gear1, temp_gear
	}
	return gear1, gear2
}

func RuneToInt(character rune) int {
	return int(character - '0')
}