// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Remove any whitespace, or hyphen charactrer and replace it with a regular whitespace character.
	// Then, split the string on whitespace characters, and trim any whitespace from each substring. As long as the substring is not empty, take the first letter of it, capitalize it, and use that as the acronym.

	 acronym := ""

	var nonAlphabeticRegex = regexp.MustCompile("[ -]")

	 var spaceSeparatedString string = nonAlphabeticRegex.ReplaceAllString(s, " ")

	 var substringArray []string = strings.Split(spaceSeparatedString, " ")

	 for index := 0; index < len(substringArray); index++ {
	 	var currentString string = strings.TrimSpace(substringArray[index])

	 	if (len(currentString) > 0){
	 		firstLetter := currentString[0]
	 		acronym += strings.ToUpper(string(firstLetter))
		}
	}

	 return acronym
}
