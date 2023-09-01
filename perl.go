package perl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
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
	perlParams := ""
	for paramName, paramValue := range p.params {
		jsonValue, err := json.Marshal(paramValue)
		if err != nil {
			return result, err
		}
		perlParams = fmt.Sprint(perlParams, "my ", "$", paramName, " = from_json(", strconv.Quote(string(jsonValue)), ");\n")
	}
	allPerlCommand := fmt.Sprintf(`
	use strict;
	use warnings;
	%s
	use JSON qw(from_json to_json);

	sub main
	{
	%s
	%s
	}
	my $result = main();
	print to_json($result);
	1;
	`, buildPerlInc(), perlParams, command)
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
