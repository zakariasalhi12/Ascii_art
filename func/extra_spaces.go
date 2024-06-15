package ascii

import "strings"

func RemoveExtraSpaces(s string) string {
	// Split the string by spaces
	parts := strings.Fields(s)

	// Join the non-empty parts with a single space
	return strings.Join(parts, " ")
}
