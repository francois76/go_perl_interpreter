package lib

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func BuildPerlparams(params map[string]interface{}) string {
	perlParams := ""
	for paramName, paramValue := range params {
		jsonValue, err := json.Marshal(paramValue)
		if err != nil {
			panic(err)
		}
		perlParams = fmt.Sprint(perlParams, "my ", "$", paramName, " = from_json(", strconv.Quote(string(jsonValue)), ");\n")
	}
	return perlParams
}
