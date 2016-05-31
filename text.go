package gadgets

import (
	"regexp"
	"strings"
)

// Slugify takes an input string, slugifies (remove special characters, replaces spaces with "-", etc.) it, and returns
// it as a string.
func Slugify(input string) string {
	// Trim the input string.
	input = strings.TrimSpace(input)

	// Make lower case.
	input = strings.ToLower(input)

	// Remove non-alphanumeric characters and hyphens.
	input = regexp.MustCompile("[^a-z0-9- ]").ReplaceAllString(input, "")

	// Replace spaces with "-".
	input = regexp.MustCompile("[\\s]+").ReplaceAllString(input, "-")

	return input
}
