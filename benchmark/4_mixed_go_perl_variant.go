package main

import (
	"fmt"
	"github.com/francois76/go-perl-interpreter"
)

func main(){

	result:= ""
	for i := 0; i < 10; i++ {
		result, _ = perl.Params[string](perl.P{
			"result":result,
			"input": i,
		}).Exec(`
	return $result .= $input;
	`)
	}
	fmt.Println(len(result))
}