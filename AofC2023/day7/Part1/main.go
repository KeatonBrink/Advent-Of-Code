package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	card string
	bid  int
}

const (
	total_cards = 13
)

const (
	c2 = iota
	c3
	c4
	c5
	c6
	c7
	c8
	c9
	cT
	cJ
	cQ
	cK
	cA
)

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

	grouped_by_win_type := make([][]Pair, 7)

	for _, line := range text {
		var cards_from_line string
		var bid_from_line int
		s, err := fmt.Sscanf(line, "%s %d", &cards_from_line, &bid_from_line)
		if err != nil {
			panic(err)
		}
		if s != 2 {
			panic("Not enough found in scan")
		}
		grouped_by_win_type[FindWinType(cards_from_line)] = append(grouped_by_win_type[FindWinType(cards_from_line)], Pair{card: cards_from_line, bid: bid_from_line})
	}

	println("Valid attempts", return_val)
}

func FindWinType(cards string) int {
	if IsFiveOfKind(cards) {
		return 6
	} else if IsFourOfKind(cards) {
		return 5
	} else if IsFullHouse(cards) {
		return 4
	} else if IsThreeOfKind(cards) {
		return 3
	} else if IsTwoPair(cards) {
		return 2
	} else if IsOnePair(cards) {
		return 1
	}
	return 0
}

func CharToInt(card byte) int {
	switch {
	case card >= '2' && card <= '9':
		return int(card - '2')
	case card == 'T':
		return 8
	case card == 'J':
		return 9
	case card == 'Q':
		return 10
	case card == 'K':
		return 11
	case card == 'A':
		return 12
	default:
		panic(fmt.Sprintf("Invalid byte %b", card))
	}
}

func IsFiveOfKind(cards string) bool {
	return cards[0] == cards[1] && cards[1] == cards[2] && cards[2] == cards[3] && cards[3] == cards[4]
}

func IsFourOfKind(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
		if potential_cards[integer_card] >= 4 {
			return true
		}
	}
	return false
}

func IsFullHouse(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}
	is_triple := false
	is_double := false
	for _, card_int := range potential_cards {
		if card_int == 3 {
			is_triple = true
		} else if card_int == 2 {
			is_double = true
		}
	}
	return is_double && is_triple
}

func IsThreeOfKind(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
		if potential_cards[integer_card] >= 3 {
			return true
		}
	}
	return false
}

func IsTwoPair(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}
	is_double := false
	for _, card_int := range potential_cards {
		if card_int == 2 {
			if !is_double {
				is_double = true
			} else {
				return true
			}
		}
	}
	return false
}

func IsOnePair(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}
	for _, card_int := range potential_cards {
		if card_int == 2 {
			return true
		}
	}
	return false
}
