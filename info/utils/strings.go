package utils

import "strings"

// Trimmer basically whitespaces and makes it lowercase. Useful for CLI flags,
// so you don't have to make hundreds of possible aliases based on spaces and
// case.
func Trimmer(str string) string {
	space := strings.ReplaceAll(str, " ", "")
	return strings.ToLower(space)
}
