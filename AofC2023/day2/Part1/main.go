package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

type GamePair struct{
	gameId int;
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

	game_config := GameConfig{blue: 14, red: 12, green: 13}

	return_channel := make(chan GamePair, 100) 

	for file_scanner.Scan() {
		total_goroutines += 1
		go scan_line(file_scanner.Text(), game_config, return_channel)
	}
	
	finished_routines := 0

	final_value := 0

	for finished_routines < total_goroutines {
		temp_ret := <- return_channel
		if temp_ret.isValid {
			final_value += temp_ret.gameId
		}
		finished_routines += 1
		println("Return value: ", temp_ret.gameId, temp_ret.isValid)
	}
	println("Final Sum: ", final_value)
}

func scan_line(curline string, game_config GameConfig, returnchan chan<- GamePair) {
	curline = curline[5:]
	// println("Curline: ", curline)
	gameid := ""
	for curline[0] != ':' {
		gameid += string(curline[0])
		curline = curline[1:]
	}
	gameidNum, err := strconv.ParseInt(gameid, 10, 64)
	if err != nil {
		panic(err)
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
			if tempVal > game_config.red {
				returnchan <- GamePair{gameId: int(gameidNum), isValid: false}
				return
			}
		case "blue":
			curline = curline[4:]
			if tempVal > game_config.blue {
				returnchan <- GamePair{gameId: int(gameidNum), isValid: false}
				return
			}
		case "green":
			curline = curline[5:]
			if tempVal > game_config.green {
				returnchan <- GamePair{gameId: int(gameidNum), isValid: false}
				return
			}
		}
		if len(curline) > 2 {
			curline = curline[2:]
		}
	}
	returnchan <- GamePair{gameId: int(gameidNum), isValid: true}
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