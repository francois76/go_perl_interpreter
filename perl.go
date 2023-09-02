package perl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/uuid"
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
		buildCustomPrinter(p.uuid),
		`

	sub main
	{
		`,
		buildPerlparams(p.params),
		sanitizeCommand(p.uuid, command),
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
	if stderr.Len() > 0 {
		return result, errors.New(stderr.String())
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Erreur lors de la création du pipe pour stdout:", err)
		return
	}

	if err = cmd.Start(); err != nil {
		fmt.Println("Erreur lors du démarrage de la commande:", err)
		return
	}

	// Créez un scanner pour lire la sortie de la commande ligne par ligne
	scanner := bufio.NewScanner(stdout)

	// Lancez une goroutine pour lire la sortie de la commande et envoyer chaque ligne dans le canal
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, "]")
			if parts[0] == "["+p.uuid+" - PRINT" {
				fmt.Println(parts[1])
			} else {
				json.Unmarshal([]byte(line), &result)
			}

		}
	}()

	// Attendez que la commande se termine
	if err := cmd.Wait(); err != nil {
		fmt.Println("La commande a échoué:", err)
	}
	return result, nil
}
