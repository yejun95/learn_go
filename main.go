package main

import (
	"fmt"
	"learn_go/mydict"
	"log"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//definition, err := dictionary.Search("first")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(definition)

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		log.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println(hello)

	// 중복 검사가 잘 되는지 확인
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		log.Println(err2)
	}
}
