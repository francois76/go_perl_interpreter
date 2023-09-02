package perl

import (
	"fmt"
	"strings"
)

func sanitizeCommand(u string, command string) string {
	printer := fmt.Sprintf("print_%s", u)
	startMatcher := []string{" ", "	", "\t", "\n", ";"}
	endMatcher := []string{" ", "	", "\t", "("}
	result := command
	for _, start := range startMatcher {
		for _, end := range endMatcher {
			result = strings.ReplaceAll(result, start+"print"+end, start+printer+end)
		}
	}

	return result
}
