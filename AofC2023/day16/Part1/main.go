package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var CYCLE_MEMORY = make(map[string][][]byte, 0)

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

	var hash_input []string

	for _, line := range text {
		strs := strings.Split(line, ",")
		hash_input = append(hash_input, strs...)
	}

	ret_val := 0

	for _, string_to_hash := range hash_input {
		ret_val += Hash(string_to_hash)
	}

	fmt.Printf("Val: %d\n", ret_val)
}

func Hash(str string) (current_value int) {
	for _, chr := range str {
		current_value += int(chr)
		current_value *= 17
		current_value %= 256
	}
	return
}
