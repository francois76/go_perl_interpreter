package main

import (
	"fmt"
)

func main(){
	result:= ""
	for i := 0; i < 100000; i++ {
		result = fmt.Sprint(result,i)
	}
	fmt.Println(len(result))
}