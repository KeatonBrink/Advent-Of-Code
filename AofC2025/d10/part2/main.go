package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	// Bits
	indicator_lights int64
	buttons          []int
	joltage          map[int]int
}

func main() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(input)
	var machines []Machine
	for _, line := range input {
		fields := strings.Fields(line)
		indicator_light_str := fields[0]
		indicator_light_str = indicator_light_str[1 : len(indicator_light_str)-1]
		indicator_light_str = strings.ReplaceAll(indicator_light_str, ".", "0")
		indicator_light_str = strings.ReplaceAll(indicator_light_str, "#", "1")
		// indicator_light, err := strconv.ParseInt(reverseString(indicator_light_str), 2, 64)
		indicator_light, err := strconv.ParseInt(indicator_light_str, 2, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		joltage_str := fields[len(fields)-1]
		joltage_str = joltage_str[1 : len(joltage_str)-1]
		joltage_slc := strings.Split(joltage_str, ",")
		joltage_map := make(map[int]int, len(indicator_light_str))
		for key, val := range joltage_slc {
			val_int, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return
			}
			joltage_map[key] = val_int
		}

		cur_machine := Machine{indicator_light, make([]int, 0), joltage_map}

		fields = fields[1 : len(fields)-1]

		for _, button := range fields {
			button = button[1 : len(button)-1]
			button_slc := strings.Split(button, ",")
			button_str := strings.Repeat("0", len(indicator_light_str))
			for _, button_char := range button_slc {
				cur_button, err := strconv.Atoi(button_char)
				if err != nil {
					fmt.Println(err)
					return
				}
				button_str = replaceByteAtIndex(button_str, '1', cur_button)
			}
			final_button, err := strconv.ParseInt(button_str, 2, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			cur_machine.buttons = append(cur_machine.buttons, int(final_button))
		}
		// fmt.Println(cur_machine.indicator_lights, cur_machine.buttons, cur_machine.joltage)
		machines = append(machines, cur_machine)
	}

	total_pushes := 0

	for _, cur_machine := range machines {
		total_buttons := len(cur_machine.buttons)
		// total_indicators := len(cur_machine.joltage)
		correct_combo_found := false
		for cnt_toggled_buttons := 1; cnt_toggled_buttons <= total_buttons && !correct_combo_found; cnt_toggled_buttons++ {
			masks := Combinations(total_buttons, cnt_toggled_buttons)
			// fmt.Println("Total Buttons, CNT_toggled_buttons", total_buttons, cnt_toggled_buttons)
			// fmt.Println(masks)
			for _, mask := range masks {
				button_indexes := BitIndexes(mask)
				// fmt.Println(mask, button_indexes)
				temp_indicator_state := 0
				for _, button_ind := range button_indexes {
					temp_indicator_state ^= cur_machine.buttons[button_ind]
				}
				// fmt.Println(temp_indicator_state, int(cur_machine.indicator_lights))
				if temp_indicator_state == int(cur_machine.indicator_lights) {
					correct_combo_found = true
					total_pushes += cnt_toggled_buttons
					// fmt.Println(total_buttons, cur_machine.indicator_lights, )
					break
				}
			}
		}
	}
	fmt.Println(total_pushes)
}

func BitIndexes(x uint) []int {
	var idx []int
	for i := 0; x != 0; i++ {
		if x&1 == 1 {
			idx = append(idx, i)
		}
		x >>= 1
	}
	return idx
}

// Stolen from online
func Combinations(n, k int) []uint {
	if k == 0 {
		return []uint{0}
	}

	var res []uint
	x := uint((1 << k) - 1) // initial k-bit number, e.g. 0011
	limit := uint(1 << n)

	for x < limit {
		res = append(res, x)

		// Gosper’s hack (next combination with k bits)
		u := x & -x
		v := x + u
		x = v + (((v ^ x) >> 2) / u)
	}
	return res
}

func replaceByteAtIndex(s string, replacement byte, index int) string {
	if index < 0 || index >= len(s) {
		return s // Index out of bounds, return original string
	}
	return s[:index] + string(replacement) + s[index+1:]
}

// reverseString takes a string as input and returns its reversed version.
func reverseString(s string) string {
	// Convert the string to a slice of runes to handle Unicode characters correctly.
	runes := []rune(s)

	// Iterate through the rune slice, swapping elements from the start and end.
	// The loop continues until the indices i and j cross each other.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the modified rune slice back to a string and return it.
	return string(runes)
}

func getInputAsLines(file_name string) ([]string, error) {
	// Read in files
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Create scanner
	file_scanner := bufio.NewScanner(f)
	file_scanner.Split(bufio.ScanLines)

	// Get lines
	var text []string
	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}
	return text, nil
}
