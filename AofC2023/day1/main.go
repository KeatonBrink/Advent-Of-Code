package main

import (
	"bufio"
	"os"
	"unicode"
)

const (
	zero  = iota //Yes
	one //Yes
	two //Yes
	three
	four //Yes
	five //Yes
	six //Yes
	seven
	eight
	nine //Yes
	nope
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

	return_int_channel := make(chan int, 100) 

	for file_scanner.Scan() {
		total_goroutines += 1
		go scan_line(file_scanner.Text(), return_int_channel)
	}
	
	finished_routines := 0

	final_value := 0

	for finished_routines < total_goroutines {
		temp_ret := <- return_int_channel
		final_value += temp_ret
		finished_routines += 1
		println("Return value: ", temp_ret)
	}
	println("Final Sum: ", final_value)
}

func scan_line(curline string, returncchan chan int) {
	first := -1
	last := -1
	for ind, char := range curline {
		temp_int := -1
		if unicode.IsDigit(char) {
			temp_int = int(char - '0')
		} else {
			test_word_int := string_ind_to_num(curline, ind)
			if test_word_int != nope {
				temp_int = test_word_int
			}
		}
		if temp_int != -1 {
			if first == -1 {
				first = temp_int
			} else {
				last = temp_int
			}
		}
	}
	var sum int
	if last == -1 {
		sum = first * 10 + first
	} else {
		sum = first * 10 + last
	}
	returncchan <- sum
}

func string_ind_to_num(curline string, ind int) int {
	var temp_sub_string string
	if len(curline) >= ind + 3{
		temp_sub_string = curline[ind:]
		temp_sub_string = temp_sub_string[:3]
		// fmt.Printf("SITM 3: %s\n", temp_sub_string)
		switch temp_sub_string {
		case "one":
			return one
		case "two":
			return two
		case "six":
			return six
		}
	}
	if len(curline) >= ind + 4 {
		temp_sub_string = curline[ind:]
		temp_sub_string = temp_sub_string[:4]
		// fmt.Printf("SITM 4: %s\n", temp_sub_string)
		switch temp_sub_string {
		case "zero":
			return zero
		case "four":
			return four
		case "five":
			return five
		case "nine":
			return nine
		}
	}
	if len(curline) >= ind + 5 {
		temp_sub_string = curline[ind:]
		temp_sub_string = temp_sub_string[:5]
		// fmt.Printf("SITM 5: %s\n", temp_sub_string)
		switch temp_sub_string {
		case "three":
			return three
		case "seven":
			return seven
		case "eight":
			return eight
		}
	}
	return nope
}