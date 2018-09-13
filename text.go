package gadgets

import (
	"regexp"
	"strings"
)

// Slugify takes an input string, slugifies (remove special characters, replaces spaces with "-", etc.) it, and returns
// it as a string.
func Slugify(input string) string {
	// Make lower case.
	input = strings.ToLower(input)

	// Deal with hyphens and underscores.
	input = strings.Replace(input, "-", " ", -1)
	input = strings.Replace(input, "_", " ", -1)

	// Remove non-alphanumeric characters.
	input = regexp.MustCompile(`[^a-z0-9 ]`).ReplaceAllString(input, "")

	// Replace spaces with "-".
	input = strings.Join(strings.Fields(input), "-")

	return input
}
