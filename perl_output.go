package perl

import (
	"fmt"
)

func (p *PerlFunction[Result]) handleOutput() {

}

func buildCustomPrinter(u string) string {
	return fmt.Sprintf(`
	sub print_%[1]s
	{
		print '[%[1]s - PRINT]' , @_;
	}
`, u)
}
