package main

import "fmt"

func main() {
	a := 2
	b := &a // b에 a 메모리 주소를 할당
	*b = 20 // pointer를 통해 b값을 바꾸면 a 값도 바뀜 -> 같은 메모리 주소를 보고있기 때문

	fmt.Println(a, *b)
}
