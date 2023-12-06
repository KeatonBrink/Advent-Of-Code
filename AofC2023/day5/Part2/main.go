package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	Destination, SD_Range int
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

	var text []string

	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}

	var seeds []int
	for _, seed := range strings.Split(text[0], " ")[1:] {
		seed1, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed1)
	}
	text = text[2:]

	// first := true

	transformers := make(map[string]map[int]Pair)

	var all_categories []string

	var current_category string

	// var next_category string

	for _, line := range text {
		// println("Line:", line)
		if len(line) < 1 {
			continue
		}
		if unicode.IsLetter(rune(line[0])) {
			split_string := strings.Split(line, "-")
			current_category = split_string[0]
			// next_category = strings.Split(split_string[2], " ")[0]
			// println(next_category)
			transformers[current_category] = make(map[int]Pair)
			// if first {
			all_categories = append(all_categories, current_category)
			// } else {
			// first = false
			// }
			// all_categories = append(all_categories, next_category)
		} else if unicode.IsDigit(rune(line[0])) {
			var source, destination, sd_range int
			found, err := fmt.Sscanf(line, "%d %d %d", &destination, &source, &sd_range)
			if err != nil {
				panic(err)
			}
			if found != 3 {
				fmt.Printf("Found only %d", found)
			}
			transformers[current_category][source] = Pair{Destination: destination, SD_Range: sd_range}
		}
	}

	closest_location := int(math.Pow(2, 60))
	// println("Closest", closest_location)

	// Loop through all the seeds
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			// Loop through the tranformer, until the end
			current_source_value := seed
			// println("\n\n\n\nSeed", seed)
			for _, transformer_key := range all_categories {
				// println("Transformer", transformer_key)
				current_transformer := transformers[transformer_key]
				for key, value := range current_transformer {
					if current_source_value >= key && current_source_value < key+value.SD_Range {
						current_source_value = value.Destination + (current_source_value - key)
						break
					}
				}
				// println("Current_soure_value", current_source_value)
			}
			if current_source_value < closest_location {
				closest_location = current_source_value
			}
		}
	}

	print("Closest Location Found:", closest_location)
}
