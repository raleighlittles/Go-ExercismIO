package brackets

import (
	"container/list"
	"errors"
)

// Implement a stack using a singly-linked list
type Stack struct {
	l *list.List
}

func (s *Stack) isEmpty() ( bool ) {
	return (s.l.Len() == 0)
}

func (s *Stack) push(char string) {
	s.l.PushBack(char)

}

func (s *Stack) pop() (char string, err error) {

	// Use type assertion, see: https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion

	back := s.l.Back()

	if (back != nil){

	result, ok := back.Value.(string)

	if ok {
		char = result
		err = nil

		// now, delete the element from the back
		s.l.Remove(back)
	}


	} else {
		char = ""
		err = errors.New("Failed to pop")
	}

	return

}

func Bracket(input string) (isValid bool) {
	// Uses the canonical stack approach to check whether a string has a balanced parenthesis set.
	// General algorithm is:
	// 1) Iterate over the string
	// 2) At each iteration, check if the current character is one of the valid starting characters
	//    (usually a left parenthesis, but can technically be anything)
	//    If it is a valid starting character, push the character onto the string
	//    If its a valid ending character instead, then pop an element from the stack and check
	//    that the element you popped is the corresponding 'ender' for the current element you're on.
	//    For example, if the character you popped was a '(', then make sure the current character is
	//    a ')'.
	//    In this implementation, this is achieved by checking the indices in the corresponding
	//    delineator arrays -- such that the element at index i in the starting array is matched with
	//    and ender at index i in the ending array.
	//
	//   If at any time, you either pop a mismatch character, or you pop from an empty array,
	// then return false.
	//
	// By the time you finish iterating over the array, the stack should be empty.


	start_delineators := []string{"(", "{", "["}
	end_delineators := []string{")", "}", "]"}

	stack := Stack{list.New()}

	for _, char := range input {

		//is_starting_character, is_ending_character := false, false

		char_as_string := string(char)

		_, err := Find(start_delineators, char_as_string)

		if (err != nil) {

			// element is not a starting sequence, check if its an ending sequence?
			index_in_enders, err := Find(end_delineators, char_as_string)

			if (err != nil) {
				// element is neither a starting sequence nor ending sequence, ignore and continue
				continue
			}

			// element must be an ending sequence here, so pop off the stack
			// and check indices of popped element versus current element

			el, err := stack.pop()

			if (err != nil) {
				// you failed to pop off the stack, so abort
				return false
			}

			index_in_starters, err := Find(start_delineators, el)

			if (err != nil) {
				// the element you popped is not a start sequence, so abort
				return false
			}

			// the element you popped is a start sequence, so make sure the indices match
			if (index_in_starters != index_in_enders) {
				return false
			}

		}  else {
			// The character you're on is a start character, so push it to the stack
			stack.push(char_as_string)
		}
	}

	return stack.isEmpty()
}

func Find(arr []string, input string) (index int, err error) {
	for index, element := range arr {
		if element == input {
			return index, nil
		}
	}

	return 0, errors.New("Element not found!")
}
