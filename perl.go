package perl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Params[Result any](params P) *PerlFunction[Result] {
	return &PerlFunction[Result]{
		params: params,
	}
}

func New[Result any]() *PerlFunction[Result] {
	return &PerlFunction[Result]{
		params: map[string]interface{}{},
	}
}

func (p *PerlFunction[Result]) Exec(command string) (result Result, err error) {
	cmd := exec.Command("perl")

	allPerlCommand := fmt.Sprint(`
	use strict;
	use warnings;`,
		buildPerlInc(),
		`
	use JSON qw(from_json to_json);

	sub main
	{
		`,
		BuildPerlparams(p),
		command,
		`
	}
	my $result = main();
	print to_json($result);
	1;
	`)
	if Debug {
		fmt.Println(allPerlCommand)
	}
	cmd.Stdin = bytes.NewBufferString(allPerlCommand)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, _ := cmd.Output()
	if stderr.Len() > 0 {
		return result, errors.New(stderr.String())
	}
	outString := string(out)
	lines := strings.Split(outString, "\n")

	fmt.Println(strings.Join(lines[:len(lines)-1], "\n"))
	json.Unmarshal([]byte(lines[len(lines)-1]), &result)
	return result, nil
}
