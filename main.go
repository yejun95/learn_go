package main

import "fmt"

func rangeLoop(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	return 1
}

func forLoop(number ...int) int {
	result := 0
	for i := 0; i < len(number); i++ {
		result += number[i]
	}

	return result
}

func main() {
	rangeLoop(1, 2, 3, 4, 5, 6)
	result := forLoop(10, 20, 30, 40, 50, 60)

	fmt.Println("for loop sum :", result)
}
