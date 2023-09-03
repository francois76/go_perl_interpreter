package main

import (
	"fmt"
	"github.com/francois76/go-perl-interpreter"
)

func main(){
	result, err := perl.New[int64]().Exec(`
	my $result = "";
	for my $i (0..99999) {
		$result .= $i;
	}
	return length($result);
	`)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
}