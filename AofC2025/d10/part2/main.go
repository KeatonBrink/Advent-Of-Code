package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	buttons [][]int
	joltage []int
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
	max_joltage := -1
	for _, line := range input {
		fields := strings.Fields(line)

		joltage_str := fields[len(fields)-1]
		joltage_str = joltage_str[1 : len(joltage_str)-1]
		joltage_slc := strings.Split(joltage_str, ",")
		joltage := make([]int, 0)
		for _, val := range joltage_slc {
			val_int, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return
			}
			if val_int > max_joltage {
				max_joltage = val_int
			}
			joltage = append(joltage, val_int)
		}

		cur_machine := Machine{make([][]int, 0), joltage}

		fields = fields[1 : len(fields)-1]

		for _, button := range fields {
			cur_holder := make([]int, len(joltage))
			button = button[1 : len(button)-1]
			button_slc := strings.Split(button, ",")
			for _, button_char := range button_slc {
				cur_joltage_affect, err := strconv.Atoi(button_char)
				if err != nil {
					fmt.Println(err)
					return
				}
				cur_holder[cur_joltage_affect] = 1
			}
			cur_machine.buttons = append(cur_machine.buttons, cur_holder)
			// fmt.Println(cur_holder)
		}
		// fmt.Println(cur_machine.indicator_lights, cur_machine.buttons, cur_machine.joltage)
		machines = append(machines, cur_machine)
	}

	total_pushes := 0

	for debug_ind, machine := range machines {
		// fmt.Println(machine.buttons)

		// coeffs := machine.buttons
		coeffs := Transpose(machine.buttons)

		// fmt.Println(coeffs)

		bounds := make([]int, len(machine.buttons))

		for ind1, button_affect := range machine.buttons {
			min := max_joltage + 1
			for ind2, val := range button_affect {
				if val == 0 {
					continue
				}
				if machine.joltage[ind2] < min {
					min = machine.joltage[ind2]
				}
			}
			bounds[ind1] = min
		}

		// fmt.Println(bounds)

		// result := SolveIntegerSystem(coeffs, machine.joltage, bounds, max_joltage*len(coeffs[0]))
		_, min_pushes, _ := SolveIntegerSystemMinSum(coeffs, machine.joltage, bounds)

		// fmt.Println(result)

		// min_pushes := max_joltage * len(coeffs[0])
		// for _, possible_answer := range result {
		// 	temp_total := 0
		// 	for _, button_push := range possible_answer {
		// 		temp_total += button_push
		// 	}
		// 	if temp_total < min_pushes {
		// 		min_pushes = temp_total
		// 	}
		// }
		total_pushes += min_pushes
		fmt.Println(debug_ind, total_pushes, min_pushes)
		// fmt.Println(debug_ind)
		// GoNumLibraryAttempt(coeffs, machine.joltage, bounds)
	}

	fmt.Println(total_pushes)
}

// func SolveIntegerSystem(coeffs [][]int, rhs []int, bounds []int, maxSolutions int) (solutions [][]int) {
// 	m := len(coeffs)
// 	if m == 0 {
// 		return nil
// 	}
// 	n := len(coeffs[0])
// 	// basic validation
// 	if len(rhs) != m || len(bounds) != n {
// 		panic(fmt.Sprint("dimension mismatch ", len(rhs), m, len(bounds), n))
// 	}

// 	cur := make([]int, n)

// 	// Precompute coefficients by column for speed: coeffs[row][col] already available.
// 	// Recursive backtracking
// 	var dfs func(col int)
// 	dfs = func(col int) {
// 		// abort early if we've collected enough solutions
// 		if maxSolutions > 0 && len(solutions) >= maxSolutions {
// 			return
// 		}
// 		if col == n {
// 			// check all equalities satisfied exactly
// 			ok := true
// 			for row := 0; row < m && ok; row++ {
// 				sum := 0
// 				for c := 0; c < n; c++ {
// 					sum += coeffs[row][c] * cur[c]
// 				}
// 				if sum != rhs[row] {
// 					ok = false
// 				}
// 			}
// 			if ok {
// 				// copy solution
// 				sol := make([]int, n)
// 				copy(sol, cur)
// 				solutions = append(solutions, sol)
// 			}
// 			return
// 		}

// 		// try all values for variable col
// 		for v := 0; v <= bounds[col]; v++ {
// 			cur[col] = v

// 			// prune: for each equation check whether RHS is achievable with remaining vars
// 			prune := false
// 			for row := 0; row < m && !prune; row++ {
// 				sumAssigned := 0
// 				for c := 0; c <= col; c++ {
// 					sumAssigned += coeffs[row][c] * cur[c]
// 				}
// 				// compute min/max possible contribution of remaining variables (col+1 .. n-1)
// 				remMin, remMax := 0, 0
// 				for c := col + 1; c < n; c++ {
// 					coef := coeffs[row][c]
// 					if coef >= 0 {
// 						remMin += 0
// 						remMax += coef * bounds[c]
// 					} else {
// 						remMin += coef * bounds[c] // negative
// 						remMax += 0
// 					}
// 				}
// 				minTotal := sumAssigned + remMin
// 				maxTotal := sumAssigned + remMax
// 				if rhs[row] < minTotal || rhs[row] > maxTotal {
// 					prune = true
// 				}
// 			}

// 			if !prune {
// 				dfs(col + 1)
// 				if maxSolutions > 0 && len(solutions) >= maxSolutions {
// 					return
// 				}
// 			}
// 		}

// 		// optional: reset cur[col] (not necessary)
// 		cur[col] = 0
// 	}

// 	dfs(0)
// 	return solutions
// }

// This attempt to mimize
func SolveIntegerSystemMinSum(
	coeffs [][]int,
	rhs []int,
	bounds []int,
) (bestSolution []int, bestSum int, ok bool) {

	m := len(coeffs)
	if m == 0 {
		return nil, 0, false
	}
	n := len(coeffs[0])
	if len(rhs) != m || len(bounds) != n {
		panic("dimension mismatch")
	}

	bounds[0]++

	cur := make([]int, n)
	bestSum = math.MaxInt

	var dfs func(col int, currentSum int)
	dfs = func(col int, currentSum int) {
		// Branch-and-bound: sum(x) already worse than best
		if currentSum >= bestSum {
			return
		}

		if col == n {
			// Check feasibility
			for row := 0; row < m; row++ {
				sum := 0
				for c := 0; c < n; c++ {
					sum += coeffs[row][c] * cur[c]
				}
				if sum != rhs[row] {
					return
				}
			}

			// Found better solution
			bestSum = currentSum
			fmt.Println("New Best Sum", bestSum)
			bestSolution = append([]int(nil), cur...)
			ok = true
			return
		}

		// Try values in increasing order (important!)
		for v := 0; v <= bounds[col]; v++ {
			cur[col] = v
			newSum := currentSum + v

			// Sum-based prune
			if newSum >= bestSum {
				break // increasing v only makes it worse
			}

			// Equation-based pruning
			prune := false
			for row := 0; row < m && !prune; row++ {
				sumAssigned := 0
				for c := 0; c <= col; c++ {
					sumAssigned += coeffs[row][c] * cur[c]
				}

				remMin, remMax := 0, 0
				for c := col + 1; c < n; c++ {
					coef := coeffs[row][c]
					if coef >= 0 {
						remMax += coef * bounds[c]
					} else {
						remMin += coef * bounds[c]
					}
				}

				minTotal := sumAssigned + remMin
				maxTotal := sumAssigned + remMax
				if rhs[row] < minTotal || rhs[row] > maxTotal {
					prune = true
				}
			}

			if !prune {
				dfs(col+1, newSum)
			}
		}

		cur[col] = 0
	}

	dfs(0, 0)
	return bestSolution, bestSum, ok
}

// func GoNumLibraryAttempt(coeffs [][]int,
// 	rhs []int,
// 	bounds []int) {
// 	var single_coeffs []float64

// 	for _, elem1 := range coeffs {
// 		for _, elem2 := range elem1 {
// 			single_coeffs = append(single_coeffs, float64(elem2))
// 		}
// 	}
// 	A := mat.NewDense(len(coeffs), len(coeffs[0]), single_coeffs)

// 	var b_float []float64

// 	for _, elem1 := range rhs {
// 		b_float = append(b_float, float64(elem1))
// 	}
// 	b := mat.NewVecDense(len(rhs), b_float)

// 	var x mat.VecDense

// 	// Solve A * x = b
// 	if err := x.SolveVec(A, b); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("x =\n%v\n", mat.Formatted(&x, mat.Prefix("  "), mat.Excerpt(0)))
// }

func Transpose[T any](m [][]T) [][]T {
	if len(m) == 0 {
		return nil
	}

	rows := len(m)
	cols := len(m[0])

	// Allocate transposed matrix
	t := make([][]T, cols)
	for i := range t {
		t[i] = make([]T, rows)
	}

	// Transpose
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			t[j][i] = m[i][j]
		}
	}

	return t
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
