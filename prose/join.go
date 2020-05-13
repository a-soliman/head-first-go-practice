package prose

import "strings"

// JoinWithCommas given a slice of strings, it returns a single string seperated with commas
func JoinWithCommas(phrase []string) string {
	var finalString string

	switch len(phrase) {
	case 0:
		finalString = ""
	case 1:
		finalString = phrase[0]
	case 2:
		finalString = phrase[0] + " and " + phrase[1]
	default:
		finalString = strings.Join(phrase[:len(phrase)-1], ", ")
		finalString += ", and "
		finalString += phrase[len(phrase)-1]
	}

	return finalString
}
