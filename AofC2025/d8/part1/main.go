package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x float64
	y float64
	z float64
}

func (jb1 JunctionBox) IsSame(jb2 JunctionBox) bool {
	return jb1.x == jb2.x && jb1.y == jb2.y && jb1.z == jb2.z
}

func (jb JunctionBox) Print() {
	fmt.Println(jb, jb.x, jb.y, jb.z)
}
func (jb JunctionBox) Stringify() string {
	return fmt.Sprint(jb.x, jb.y, jb.z)
}

type JunctionBoxConnection struct {
	a *JunctionBox
	b *JunctionBox
}

func (jbc JunctionBoxConnection) Distance() float64 {
	x_squared := math.Pow(jbc.a.x-jbc.b.x, 2)
	y_squared := math.Pow(jbc.a.y-jbc.b.y, 2)
	z_squared := math.Pow(jbc.a.z-jbc.b.z, 2)
	return math.Pow(x_squared+y_squared+z_squared, 0.5)
}

func (jbc JunctionBoxConnection) Print() {
	fmt.Println(jbc.a.Stringify(), ", ", jbc.b.Stringify(), ":::", jbc.Distance())
}

func main() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	max_pairs := 10
	TOTAL_MAXES := 3
	if file_name == "input.txt" {
		max_pairs = 1000
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	var junction_boxes []JunctionBox
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		junction_boxes = append(junction_boxes, JunctionBox{
			float64(x),
			float64(y),
			float64(z)})
	}
	fmt.Println("Junction Boxes Created")
	var junction_pairs []JunctionBoxConnection
	for ind1 := 0; ind1 < len(junction_boxes)-1; ind1++ {
		for ind2 := ind1 + 1; ind2 < len(junction_boxes); ind2++ {
			// fmt.Println(ind1, ind2)
			jbc := JunctionBoxConnection{&junction_boxes[ind1], &junction_boxes[ind2]}
			junction_pairs = append(junction_pairs, jbc)
		}
	}
	sort.Slice(junction_pairs, func(i, j int) bool {
		return junction_pairs[i].Distance() < junction_pairs[j].Distance()
	})
	fmt.Println("Junction Box Connections completed")
	pairs := 0
	var circuits [][]*JunctionBox
	for _, jbc := range junction_pairs {
		// fmt.Println(jbc.a, jbc.b)
		if pairs == max_pairs {
			break
		}
		a_ind := -1
		b_ind := -1
		for ind, circuit := range circuits {
			for _, junction_box := range circuit {
				// Probably could have compared address, but this works as well
				if junction_box.IsSame(*jbc.a) {
					a_ind = ind
				}
				if junction_box.IsSame(*jbc.b) {
					b_ind = ind
				}
			}
		}
		if a_ind == -1 {
			if b_ind == -1 {
				circuits = append(circuits, []*JunctionBox{jbc.a, jbc.b})
			} else if b_ind != -1 {
				circuits[b_ind] = append(circuits[b_ind], jbc.a)
			}
		} else if a_ind != -1 {
			if b_ind == -1 {
				circuits[a_ind] = append(circuits[a_ind], jbc.b)
			} else if b_ind != -1 && a_ind != b_ind {
				// Two circuits need to be combined
				circuits[a_ind] = append(circuits[a_ind], circuits[b_ind]...)
				circuits = append(circuits[:b_ind], circuits[b_ind+1:]...)
			}
		}
		pairs++
	}
	fmt.Println("Circuits Created")

	var maxes []int
	for i := 0; i < TOTAL_MAXES; i++ {
		maxes = append(maxes, -1)
	}

	for _, circuit := range circuits {
		circuit_size := len(circuit)
		for ind1, cur_max := range maxes {
			if cur_max == -1 {
				maxes[ind1] = circuit_size
				break
			} else if cur_max < circuit_size {
				maxes = append(maxes[:ind1], append([]int{circuit_size}, maxes[ind1:]...)...)
				maxes = maxes[:TOTAL_MAXES]
				break
			}
		}
	}
	fmt.Println(maxes[0] * maxes[1] * maxes[2])

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
