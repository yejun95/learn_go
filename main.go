package main

import "fmt"

func main() {
	names := map[string]int{
		"name": 421,
		"age":  12,
	}

	for key, value := range names {
		fmt.Println("key : ", key)
		fmt.Println("value : ", value)
	}

	fmt.Println(names)
}
