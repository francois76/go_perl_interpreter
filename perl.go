package perl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/francois76/go-perl-interpreter/internal/lib"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func New[Result any]() *PerlFunction[Result] {
	return Params[Result](P{})
}

func Params[Result any](params P) *PerlFunction[Result] {
	return &PerlFunction[Result]{
		params: params,
		uuid:   strings.ReplaceAll(uuid.New().String(), "-", ""),
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

	`,
		lib.BuildCustomPrinter(p.uuid),
		`

	sub main_`, p.uuid, `
	{
		`,
		lib.BuildPerlparams(p.params),
		lib.SanitizeCommand(p.uuid, command),
		`
	}
	my $result = main_`, p.uuid, `();
	print to_json($result);
	1;
	`)

	log.Debug(allPerlCommand)
	cmd.Stdin = bytes.NewBufferString(allPerlCommand)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if stderr.Len() > 0 {
		return result, errors.New(stderr.String())
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)

	// locker is here to prevent function to return before all lines are processed
	locker := make(chan bool, 1)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, "]")
			if parts[0] == "["+p.uuid+" - PRINT" {
				fmt.Println(parts[1])
			} else {
				log.Debug("returned value : " + line)
				json.Unmarshal([]byte(line), &result)
			}
		}
		locker <- true
	}()

	if err = cmd.Wait(); err != nil {
		return
	}
	<-locker
	return result, nil
}
