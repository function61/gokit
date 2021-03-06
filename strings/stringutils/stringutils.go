// Things that missing from Go's "strings" package
package stringutils

// https://stackoverflow.com/questions/28718682/how-to-get-a-substring-from-a-string-of-runes-in-golang/56129287#56129287
// https://stackoverflow.com/questions/12311033/extracting-substrings-in-go/56129336#56129336
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func Truncate(input string, to int) string {
	// we could optimize and do len(input) here but it operates on bytes (not runes), so
	// we could have some unexpected behaviour with input getting suffixed ... when no
	// truncation actually happens
	truncated := Substr(input, 0, to)

	if truncated != input {
		indicator := ".."

		return Substr(truncated, 0, max(to-len(indicator), 0)) + indicator
	}

	return input
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
