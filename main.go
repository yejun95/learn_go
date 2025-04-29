package main

import (
	"fmt"
	"learn_go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)

}
