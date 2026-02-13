package utils

import "regexp"

// split at "|" with any amount of spaces before and after
func SplitAtPipe(s string) []string {
	re := regexp.MustCompile(`\s*\|\s*`)
	return re.Split(s, -1)
}
