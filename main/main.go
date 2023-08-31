package main

import (
	"fmt"

	"github.com/francois76/go-perl-interpreter"
)

func main() {
	err := perl.Exec(`
	use strict;
	use warnings;
	print 10."\n";
	print 20;
	1;
	`)
	if err != nil {
		fmt.Println("test")
		fmt.Println(err)
	}
}
