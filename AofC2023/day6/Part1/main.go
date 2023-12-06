package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Destination, SD_Range int
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

	times := strings.Fields(text[0])[1:]

	distances := strings.Fields(text[1])[1:]

	return_val := 1

	for i := 0; i < len(times); i++ {
		valid_attempts := 0
		max_time, _ := strconv.Atoi(times[i])
		min_distance, _ := strconv.Atoi(distances[i])
		for seconds := 1; seconds < int(max_time); seconds++ {
			// println(seconds * (int(max_time) - seconds))
			if min_distance < seconds*(int(max_time)-seconds) {
				valid_attempts += 1
			}
		}
		return_val *= valid_attempts
	}
	println("Valid attempts", return_val)
}
