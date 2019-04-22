package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) (string) {

	encodedArray := Construct(message, rails)

	var railString string

	var encodedString string

	for i := 0; i < len(encodedArray); i++ {
		railString = strings.Join(encodedArray[i], "")
		fmt.Println("current rail =", railString)
		encodedString += railString
	}

	return encodedString

}

func Decode(message string, rails int) (string) {


}

//func Apply(matrix [][]string ) {
//	// Returns the "applied" version of the matrix.
//
//}

func Construct(message string, rails int) ([][]string) {

	matrix := make([][]string, rails)
	var outerIndex int = 0
	var forwards bool = true
	var spare int = 0

	var letters int = 0

	for letters != len(message) {

		if forwards {
			outerIndex = 0 + spare
			for outerIndex < rails  && letters < len(message) {
				fmt.Println("Current letter =", string(message[letters]), "at index", letters)
				fmt.Println("Placing current letter into row #=", outerIndex)
				matrix[outerIndex] = append(matrix[outerIndex], string(message[letters]))
				outerIndex++
				letters++


				fmt.Println("Current rails look like...")

				 for i := 0; i < len(matrix); i++ {
					fmt.Print(matrix[i])
					fmt.Print("|")
				}
			}

		} else {
			outerIndex = rails - 1 - spare
			for outerIndex > -1 && letters < len(message) {
				fmt.Println("Current letter =", string(message[letters]), "at index", letters)
				fmt.Println("Placing current letter at row #=", outerIndex)
				matrix[outerIndex] = append(matrix[outerIndex], string(message[letters]))
				outerIndex--
				letters++

				fmt.Println("Current rails look like...")

				for i := 0; i < len(matrix); i++ {
					fmt.Print(matrix[i])
					fmt.Print("|")
				}
			}
		}

		fmt.Println("Switching direction")
		forwards = !forwards
		spare = 1


	}

	return matrix
}


