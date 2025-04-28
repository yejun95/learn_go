package main

import "fmt"

func main() {
	names := []string{"aa", "bb", "cc"}
	names = append(names, "dd")
	fmt.Println(names)
}
