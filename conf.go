package perl

import "fmt"

var PerlInc = []string{}

func buildPerlInc() (result string) {
	for _, line := range PerlInc {
		result = fmt.Sprint("use lib ", line, ";\n")
	}
	return result
}
