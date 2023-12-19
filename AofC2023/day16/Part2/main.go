package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Up = iota
	Down
	Left
	Right
)

type EndOfLight struct {
	Row, Col, InputDirection int
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

	mirror_layout := text

	var visited_positions = make([][]bool, len(mirror_layout))

	var visited_EOLS []EndOfLight

	for i := 0; i < len(mirror_layout); i++ {
		visited_positions[i] = make([]bool, len(mirror_layout[i]))
	}

	var eol_positions []EndOfLight

	start_EOL := EndOfLight{Row: 0, Col: 0, InputDirection: Left}

	eol_positions = append(eol_positions, start_EOL)

	for len(eol_positions) > 0 {
		var cur_EOL EndOfLight
		cur_EOL, eol_positions = eol_positions[0], eol_positions[1:]
		if cur_EOL.Col < 0 || cur_EOL.Col >= len(mirror_layout[0]) || cur_EOL.Row < 0 || cur_EOL.Col >= len(mirror_layout) {
			println("missing edge case catch")
			continue
		}
		found_dup := false
		for _, eol := range visited_EOLS {
			if eol.Row == cur_EOL.Row && cur_EOL.Col == eol.Col && cur_EOL.InputDirection == eol.InputDirection {
				found_dup = true
				break
			}
		}
		if found_dup {
			continue
		}
		visited_EOLS = append(visited_EOLS, cur_EOL)
		PrintEOL(cur_EOL)
		visited_positions[cur_EOL.Row][cur_EOL.Col] = true
		cur_apparatus := mirror_layout[cur_EOL.Row][cur_EOL.Col]
		switch cur_EOL.InputDirection {
		case Up:
			if cur_EOL.Col < len(mirror_layout[cur_EOL.Row])-1 && (cur_apparatus == '\\' || cur_apparatus == '-') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row, Col: cur_EOL.Col + 1, InputDirection: Left})
			}
			if cur_EOL.Row < len(mirror_layout)-1 && (cur_apparatus == '.' || cur_apparatus == '|') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row + 1, Col: cur_EOL.Col, InputDirection: Up})
			}
			if cur_EOL.Col > 0 && (cur_apparatus == '/' || cur_apparatus == '-') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row, Col: cur_EOL.Col - 1, InputDirection: Right})
			}
		case Down:
			if cur_EOL.Col < len(mirror_layout[cur_EOL.Row])-1 && (cur_apparatus == '/' || cur_apparatus == '-') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row, Col: cur_EOL.Col + 1, InputDirection: Left})
			}
			if cur_EOL.Row > 0 && (cur_apparatus == '.' || cur_apparatus == '|') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row - 1, Col: cur_EOL.Col, InputDirection: Down})
			}
			if cur_EOL.Col > 0 && (cur_apparatus == '\\' || cur_apparatus == '-') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row, Col: cur_EOL.Col - 1, InputDirection: Right})
			}
		case Left:
			if cur_EOL.Col < len(mirror_layout[cur_EOL.Row])-1 && (cur_apparatus == '.' || cur_apparatus == '-') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row, Col: cur_EOL.Col + 1, InputDirection: Left})
			}
			if cur_EOL.Row < len(mirror_layout)-1 && (cur_apparatus == '\\' || cur_apparatus == '|') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row + 1, Col: cur_EOL.Col, InputDirection: Up})
			}
			if cur_EOL.Row > 0 && (cur_apparatus == '/' || cur_apparatus == '|') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row - 1, Col: cur_EOL.Col, InputDirection: Down})
			}
		case Right:
			if cur_EOL.Col > 0 && (cur_apparatus == '.' || cur_apparatus == '-') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row, Col: cur_EOL.Col - 1, InputDirection: Right})
			}
			if cur_EOL.Row < len(mirror_layout)-1 && (cur_apparatus == '/' || cur_apparatus == '|') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row + 1, Col: cur_EOL.Col, InputDirection: Up})
			}
			if cur_EOL.Row > 0 && (cur_apparatus == '\\' || cur_apparatus == '|') {
				eol_positions = append(eol_positions, EndOfLight{Row: cur_EOL.Row - 1, Col: cur_EOL.Col, InputDirection: Down})
			}
		}
	}

	ret_val := 0

	for _, row := range visited_positions {
		for _, val := range row {
			if val == true {
				ret_val++
			}
		}
	}

	fmt.Printf("Final lens hash: %d\n", ret_val)
}

func PrintEOL(eol EndOfLight) {
	fmt.Printf("Row %d Col %d InputDirection %d\n", eol.Row, eol.Col, eol.InputDirection)
}
