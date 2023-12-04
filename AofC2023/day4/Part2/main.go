package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	GameId, Matches int
}

func main() {
	// input_file_name := "input.txt"
	input_file_name := "test_input.txt"

	read_file, err := os.Open(input_file_name)
	if err != nil {
		panic(err)
	}

	file_scanner := bufio.NewScanner(read_file)
	file_scanner.Split(bufio.ScanLines)

	total_goroutines := 0

	return_channel := make(chan Pair, 100)

	for file_scanner.Scan() {
		go scan_line(file_scanner.Text(), total_goroutines, return_channel)
		total_goroutines += 1
	}

	finished_routines := 0

	return_values := make([]int, total_goroutines)

	for finished_routines < total_goroutines {
		temp_ret := <-return_channel
		return_values[temp_ret.GameId] = temp_ret.Matches
		finished_routines += 1
	}

	card_copies := make([]int, total_goroutines)

	for i := range card_copies {
		card_copies[i] = 1
	}

	total_scratchcards := 0

	for i := 0; i < total_goroutines; i++ {
		total_scratchcards += card_copies[i]
		println(total_scratchcards)
		for j := 1; j+i < total_goroutines && j <= card_copies[i]; j++ {
			card_copies[i+j] += 1
		}
	}

	println("Final Sum: ", total_scratchcards-1)
}

func scan_line(curline string, gameid int, returnchan chan<- Pair) {
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
		for _, val := range winning_numbers {
			if temp_int == val {
				matches += 1
				break
			}
		}
		if len(curline) == 0 {
			break
		}
		for curline[0] == ' ' {
			curline = curline[1:]
		}
	}

	returnchan <- Pair{
		GameId: gameid, Matches: matches,
	}
}
