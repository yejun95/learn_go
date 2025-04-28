package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("function이 끝났을 때 실행")
	fmt.Println("------------------")
	defer fmt.Println("111111")
	fmt.Println(";;;;;;;;;;;;;;;;;;;")
	defer fmt.Println("2222222222") // 제일 마지막에 있는 defer 부터 function이 끝나고 실행된다.
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)
}
