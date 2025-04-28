package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("function이 끝났을 때 실행")
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)
}
