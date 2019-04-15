package queenattack

import (
	"errors"
	"strconv"
)

type Position struct {
	row int
	column int
}

func CanQueenAttack(white string, black string)  (attack bool, ok error) {
	if (white == black) {
		return false, errors.New("Same square")
}

	if len(white) != 2 || len(black) != 2 {
		return false, errors.New("Invalid")
	}

	whitePosition, err := ConvertStringToPosition(white)
	if (err != nil) {
		return false, errors.New("Couldn't convert string to actual coordinates")
}

	blackPosition, err := ConvertStringToPosition(black)

	if (err != nil) {
		return false, errors.New("Couldn't convert string to actual coordinates")
}

	// Most obvious case, are they in the same row or column?
	if (whitePosition.row == blackPosition.row ||
		whitePosition.column == blackPosition.column) {
		return true, nil
	}

	// More tricky case: coordinates are diagonal adjacent.
	if Abs(whitePosition.row - blackPosition.row) == Abs(whitePosition.column - blackPosition.column) {
		return true, nil
	}

	// No attack.
	return false, nil

}

// Again, Golang let me down by not having a Math.Abs function for int64 types.

func Abs (x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func ConvertStringToPosition(position string) (Position, error) {
	columnNumber, err := strconv.Atoi(string(position[1]))

	if (err != nil) {
		panic(1)
	}

	if (columnNumber < 1 || columnNumber > 8) {
		return Position{}, errors.New("Invalid column")
	}

	var rowLetter string = string(position[0])
	rowLetterMap := map[string]int{"a" : 1, "b": 2, "c":3, "d":4, "e":5,
						"f" : 6, "g": 7, "h": 8}


	var rowNumber int
	if value, ok := rowLetterMap[rowLetter]; ok {
		rowNumber = value
	} else {
		return Position{}, errors.New("Invalid row")
	}

	return Position{rowNumber, columnNumber}, nil

}
