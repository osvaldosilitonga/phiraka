package helpers

import "strings"

func SplitError(err error) string {
	splitErr := strings.TrimSpace(strings.Split(err.Error(), "Error:")[1])
	return strings.Split(splitErr, "\n")[0]
}
