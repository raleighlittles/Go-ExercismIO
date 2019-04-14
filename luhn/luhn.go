package luhn

import (
	"regexp"
	"strconv"
)

func Valid(input string) bool {

	if len(input) <= 1 {
		return false
	}


	var whitespaceRegex = regexp.MustCompile("\\s")

	space_separated_input := whitespaceRegex.ReplaceAllString(input, "")

	match, _ := regexp.MatchString("[a-zA-Z]", space_separated_input)

	if (match) {
		return false
	}

	stringLength := len(space_separated_input)

	if (stringLength <= 1) {
		return false
	}

	doubled_string := ""

	for position, character := range space_separated_input {
			// If the length of the string is even, then the positions you are doubling
			// will be 0, 2, 4, ... (which is x % 2 = 0)

		// If the length of the string is odd, then the positions you are doubling
		// will be 1,3,5, ... (which is x % 2 = 1)

		expression := (stringLength % 2 == 0)
		modulus := 0
		 if (expression) {
				modulus = 0
		 } else {
			 modulus = 1

		 }

			if position % 2 == modulus {
				correctedDigitString := ""
				currentDigitInteger, _ := strconv.Atoi(string(character))

				correctedDigitString = strconv.Itoa(returnCorrectedNumber(currentDigitInteger))

				doubled_string += correctedDigitString
			} else {
				doubled_string += string(character)
			}
		}



	digitSum := 0

	for _, character := range doubled_string {

		currentDigit, _ := strconv.Atoi(string(character))

		digitSum += currentDigit
	}

	return (digitSum % 10 == 0)

}

func returnCorrectedNumber(input int) int {
	var doubled_number = input * 2

	if doubled_number > 9 {
		return doubled_number - 9
	} else {
		return doubled_number
	}

}
