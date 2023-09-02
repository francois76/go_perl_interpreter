package perl

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func buildPerlparams[Result any](p *PerlFunction[Result]) string {
	perlParams := ""
	for paramName, paramValue := range p.params {
		jsonValue, err := json.Marshal(paramValue)
		if err != nil {
			panic(err)
		}
		perlParams = fmt.Sprint(perlParams, "my ", "$", paramName, " = from_json(", strconv.Quote(string(jsonValue)), ");\n")
	}
	return perlParams
}
