package day3

import (
	"log"
	"math"
	"os"
)

func ParseFile(filename string) int {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ParseString(string(f))
}

func ParseString(input string) int {
	finalMult := 0
	for i, ch := range input {
		if ch == rune('m') && i+2 < len(input) {
			if int32(input[i+1]) == rune('u') && int32(input[i+2]) == rune('l') {
				// mul is recognized
				left, ok, indexOfComma := readUntilComma(input[i+2:])
				if !ok {
					continue
				}
				// +4 because we pass input[i+2] and at the time of taking the index
				// inside readUntilBracket we loop through input[2:] (2+2=4)
				right, ok := readUntilBracket(input[i+indexOfComma+4:])
				if !ok {
					continue
				}
				finalMult += left * right

			}
		}
	}
	return finalMult
}

func readUntilComma(line string) (int, bool, int) {
	var finalNumber []rune
	if int32(line[1]) != rune('(') {
		return 0, false, 0
	}
	for i, ch := range line[2:] {
		if isNumber(ch) {
			finalNumber = append(finalNumber, ch)
		} else if ch == rune(',') {
			return toNumber(finalNumber), true, i
		} else {
			return 0, false, 0
		}
	}
	return 0, false, 0
}

func isNumber(ch rune) bool {
	return ch >= rune('0') && ch <= rune('9')
}

func toNumber(num []rune) (finalNumber int) {
	for i := 0; i < len(num); i++ {
		// ['4','2','0']
		finalNumber += (int(num[i]) - '0') * int(math.Pow(float64(10), float64(len(num)-i-1)))
	}
	return
}

func readUntilBracket(line string) (int, bool) {
	var finalNumber []rune
	for _, ch := range line[1:] {
		if isNumber(ch) {
			finalNumber = append(finalNumber, ch)
		} else if ch == rune(')') {
			return toNumber(finalNumber), true
		} else {
			return 0, false
		}
	}
	return 0, false
}
