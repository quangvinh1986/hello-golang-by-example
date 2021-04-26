package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(calcSum(10))
	var number int
	number = 10
	fmt.Println(number)
}

func calcSum(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}
