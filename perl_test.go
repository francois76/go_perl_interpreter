package perl_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/francois76/go-perl-interpreter"
	"github.com/maxatome/go-testdeep/td"
)

type sample struct {
	AString string
	AInt    int
	ABool   perl.Bool
}

func TestPackage(t *testing.T) {

	t.Run("success_input", func(t *testing.T) {
		result, err := successInput(sample{
			AString: "test",
			AInt:    52,
			ABool:   true,
		}, 2)

		td.CmpNoError(t, err)
		td.Cmp(t, result, perl.Void(0))
	})

	t.Run("success_output", func(t *testing.T) {
		result, err := perl.New[sample]().Exec(`
		return {
			AString => 'myResult',
			AInt => 96,
			ABool => 1
		}
		`)
		td.CmpNoError(t, err)
		td.Cmp(t, result, sample{
			AString: "myResult",
			AInt:    96,
			ABool:   true,
		})
	})

	t.Run("success_call_script", func(t *testing.T) {
		currentDir, err := os.Getwd()
		if err != nil {
			t.Fail()
		}
		result, err := perl.New[string]().Exec(fmt.Sprintf(`
		use lib "%s/";
		use testPerlScript;
		return testPerlScript::hello();
		`, currentDir))
		td.CmpNoError(t, err)
		td.Cmp(t, result, "hello perl")
	})
}

// successInput
func successInput(firstKey sample, secondKey int) (perl.Void, error) {
	return perl.Params[perl.Void](perl.P{
		"firstKey": sample{
			AString: "test",
			AInt:    52,
			ABool:   true,
		},
		"secondKey": 2,
	}).Exec(`
print "$firstKey->{AString} \n";
print "$firstKey->{AInt} \n";
print "$firstKey->{ABool} \n";
print "$secondKey\n";
return;
`)
}
