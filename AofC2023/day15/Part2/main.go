package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Name   string
	Length int
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

	var hash_input []string

	for _, line := range text {
		strs := strings.Split(line, ",")
		hash_input = append(hash_input, strs...)
	}

	lens_box_hashmap := make(map[int][]Lens, 0)

	for _, composite_entry_string := range hash_input {

		if strings.Index(composite_entry_string, "-") == -1 {
			temp_split := strings.Split(composite_entry_string, "=")
			if len(temp_split) != 2 {
				panic("= not found")
			}
			var err error
			var lens_length int
			lens_length, err = strconv.Atoi(temp_split[1])
			if err != nil {
				panic(err)
			}
			cur_lens := Lens{Name: temp_split[0], Length: lens_length}
			lens_hash := Hash(cur_lens.Name)
			var box []Lens
			var ok bool
			box, ok = lens_box_hashmap[lens_hash]
			if ok {
				found := false
				for ind, lens_in_box := range box {
					if lens_in_box.Name == cur_lens.Name {
						lens_box_hashmap[lens_hash][ind] = cur_lens
						found = true
						break
					}
				}
				if !found {
					lens_box_hashmap[lens_hash] = append(lens_box_hashmap[lens_hash], cur_lens)
				}
			} else {
				var new_box []Lens
				new_box = append(new_box, cur_lens)
				lens_box_hashmap[lens_hash] = new_box
			}
		} else {
			composite_entry_string = composite_entry_string[:len(composite_entry_string)-1]
			removal_hash := Hash(composite_entry_string)
			box, ok := lens_box_hashmap[removal_hash]
			if ok {
				for ind, lens_in_box := range box {
					if lens_in_box.Name == composite_entry_string {
						if ind == len(box)-1 {
							lens_box_hashmap[removal_hash] = box[:len(box)-1]
						} else {
							lens_box_hashmap[removal_hash] = append(box[:ind], box[ind+1:]...)
						}
					}
				}
			}
		}
	}

	ret_val := 0
	for _, box := range lens_box_hashmap {
		for pos_lens_in_box, lens := range box {
			box := Hash(lens.Name) + 1
			slot := pos_lens_in_box + 1
			length := lens.Length
			fmt.Printf("Name %s, Box %d, Slot %d, Length %d\n", lens.Name, box, slot, length)
			ret_val += box * slot * length
		}
	}

	fmt.Printf("Final lens hash: %d\n", ret_val)
}

func Hash(str string) (current_value int) {
	for _, chr := range str {
		current_value += int(chr)
		current_value *= 17
		current_value %= 256
	}
	return
}
