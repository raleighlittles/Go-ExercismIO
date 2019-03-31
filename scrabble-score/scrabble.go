package scrabble

import (
	"strings"
)

func Score(word string) int {
	/* Scrabble scoring rules:

		Letter                           Value
	A, E, I, O, U, L, N, R, S, T       1
	D, G                               2
	B, C, M, P                         3
	F, H, V, W, Y                      4
	K                                  5
	J, X                               8
	Q, Z                               10

	*/
	scrabbleScore := 0

	// From what I can tell to Unicode encoding, the "character" type here is actually
	// represented as a one-byte integer representing the unicode sequence and not
	// a regular string type.
	for _, character := range strings.ToUpper(word) {
		characterString := string(character)
		letterScore := 0
		if characterString == "K" {
			letterScore = 5
		} else if characterString == "J" || characterString == "X" {
			letterScore = 8
		} else if characterString == "Q" || characterString == "Z" {
			letterScore = 10
		} else if contains([]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}, characterString) {
			letterScore = 1
		} else if contains([]string{"D", "G"}, characterString) {
			letterScore = 2
		} else if contains([]string{"B", "C", "M", "P"}, characterString) {
			letterScore = 3
		} else if contains([]string{"F", "H", "V", "W", "Y"}, characterString) {
			letterScore = 4
		}

		scrabbleScore += letterScore
	}

	return scrabbleScore
}

/*
 * Helper function to check if a slice contains a string. It blows my mind that this is not part of a library somewhere -- or maybe it is and I just couldn't find it?
 */
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
