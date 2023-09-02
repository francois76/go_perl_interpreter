package perl

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		data := `
		print "$firstKey->{AString} \n";
		print($firstKey->{AInt});
		 print "$firstKey->{ABool} \n";
		 print "$firstKey->{AInt} \n";print "$firstKey->{ABool} \n";
	print "$secondKey\n";
		`
		result := sanitizeCommand("09a991ab-244c-43c0-9f1a-1895ee8f6934", data)
		td.CmpString(t, result, `
		print_09a991ab-244c-43c0-9f1a-1895ee8f6934 "$firstKey->{AString} \n";
		print_09a991ab-244c-43c0-9f1a-1895ee8f6934($firstKey->{AInt});
		 print_09a991ab-244c-43c0-9f1a-1895ee8f6934 "$firstKey->{ABool} \n";
		 print_09a991ab-244c-43c0-9f1a-1895ee8f6934 "$firstKey->{AInt} \n";print_09a991ab-244c-43c0-9f1a-1895ee8f6934 "$firstKey->{ABool} \n";
	print_09a991ab-244c-43c0-9f1a-1895ee8f6934 "$secondKey\n";
		`)
	})

}
