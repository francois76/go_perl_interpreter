package lib

import (
	"fmt"
)

func handleOutput() {

}

func BuildCustomPrinter(u string) string {
	return fmt.Sprintf(`
	sub print_%[1]s
	{
		print '[%[1]s - PRINT]' , @_;
	}
`, u)
}
