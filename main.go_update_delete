package main

import (
	"fmt"
	"learn_go/mydict"
)

func main() {
	/* 생성자 */
	dictionary := mydict.Dictionary{"first": "First word"}
	//definition, err := dictionary.Search("first")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(definition)

	/* Add */
	//word := "hello"
	//definition := "Greeting"
	//
	//err := dictionary.Add(word, definition)
	//if err != nil {
	//	log.Println(err)
	//}
	//hello, err := dictionary.Search(word)
	//fmt.Println(hello)
	//
	//// 중복 검사가 잘 되는지 확인
	//err2 := dictionary.Add(word, definition)
	//if err2 != nil {
	//	log.Println(err2)
	//}

	/* Update */
	//baseWord := "hello"
	//dictionary.Add(baseWord, "First")
	//err := dictionary.Update(baseWord, "Second")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//word, _ := dictionary.Search(baseWord)
	//fmt.Println(word)

	/* Delete */
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
