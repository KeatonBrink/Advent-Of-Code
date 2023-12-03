package main

import (
	"bufio"
	"os"
	"unicode"
)

type GamePair struct{
	gamePower int;
	isValid bool;
}

type GameConfig struct{
	blue, red, green int;
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

	total_goroutines := 0

	return_channel := make(chan GamePair, 100) 

	for file_scanner.Scan() {
		total_goroutines += 1
		go scan_line(file_scanner.Text(), return_channel)
	}
	
	finished_routines := 0

	final_value := 0

	for finished_routines < total_goroutines {
		temp_ret := <- return_channel
		if temp_ret.isValid {
			final_value += temp_ret.gamePower
		}
		finished_routines += 1
		// println("Return value: ", temp_ret.gamePower, temp_ret.isValid)
	}
	println("Final Sum: ", final_value)
}

func scan_line(curline string, returnchan chan<- GamePair) {
	gameMins := GameConfig{}
	for curline[0] != ':' {
		curline = curline[1:]
	}
	curline = curline [2:]
	for len(curline) > 0 {
		tempVal := 0
		for unicode.IsDigit(rune(curline[0])) {
			if tempVal > 0 {
				tempVal *= 10
			}
			tempVal += int(curline[0] - '0')
			curline = curline[1:]
		}
		curline = curline[1:]
		switch parse_color(curline) {
		case "red": 
			curline = curline[3:]
			if tempVal > gameMins.red {
				gameMins.red = tempVal
			}
		case "blue":
			curline = curline[4:]
			if tempVal > gameMins.blue {
				gameMins.blue = tempVal
			}
		case "green":
			curline = curline[5:]
			if tempVal > gameMins.green {
				gameMins.green = tempVal
			}
		}
		if len(curline) > 2 {
			curline = curline[2:]
		}
	}
	returnchan <- GamePair{gamePower: gameMins.blue * gameMins.green * gameMins.red, isValid: true}
}

func parse_color(curline string) string {
	// print("In parse color:", curline, "\n")
	if len(curline) >= 3 {
		if curline[:3] == "red" {
			return "red"
		}
	}
	if len(curline) >= 4 {
		if curline[:4] == "blue" {
			return "blue"
		}
	}
	if len(curline) >= 5 {
		if curline[:5] == "green" {
			return "green"
		}
	}
	println("Failure:", curline[:3], curline[:4], curline[:5])
	panic("PARSE COLOR FAILED")
}