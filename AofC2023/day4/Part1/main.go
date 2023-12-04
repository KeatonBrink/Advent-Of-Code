package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
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

	total_goroutines := 0


	return_channel := make(chan int, 100) 

	for file_scanner.Scan() {
		total_goroutines += 1
		go scan_line(file_scanner.Text(), return_channel)
	}
	
	finished_routines := 0

	final_value := 0

	for finished_routines < total_goroutines {
		temp_ret := <- return_channel
		final_value += temp_ret
		finished_routines += 1
	}
	println("Final Sum: ", final_value)
}

func scan_line(curline string, returnchan chan<- int) {
	end_of_starter := strings.Index(curline, ":")
	curline = curline[end_of_starter:]

	for !unicode.IsDigit(rune(curline[0])) {
		curline = curline[1:]
	}

	// println(curline)

	var winning_numbers []int
	for curline[0] != '|' {
		temp_int_string := ""
		for unicode.IsDigit(rune(curline[0])) {
			temp_int_string += string(curline[0])
			curline = curline[1:]
		}
		temp_int, err := strconv.Atoi(temp_int_string)
		if err != nil {
			panic(err)
		}
		winning_numbers = append(winning_numbers, temp_int)
		for curline[0] == ' ' {
			curline = curline[1:]
		}
	}

	for !unicode.IsDigit(rune(curline[0])) {
		curline = curline[1:]
	}

	matches := 0

	for len(curline) > 0 {
		temp_int_string := ""
		for unicode.IsDigit(rune(curline[0])) {
			temp_int_string += string(curline[0])
			if len(curline) <= 1 {
				curline = ""
				break
			}
			curline = curline[1:]
		}
		temp_int, err := strconv.Atoi(temp_int_string)
		if err != nil {
			println(curline)
			panic(err)
		}
		for _, val := range(winning_numbers) {
			if temp_int == val {
				matches += 1
				break
			}
		}
		if (len(curline) == 0) {
			break
		}
		for curline[0] == ' ' {
			curline = curline[1:]
		}
	}

	if matches > 0 {
		returnchan <- int(math.Pow(2.0, float64(matches - 1)))
	} else {
		returnchan <- 0
	}
}