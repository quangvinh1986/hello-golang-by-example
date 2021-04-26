package main

import "fmt"

func main() {
	number := 100
	sumAll := getSumOfNumber(number)
	fmt.Println("sum of all number from 1 to", number, "is: ", sumAll)
}

func getSumOfNumber(number int) int {
	sum := 0
	for index := 0; index <= number; index++ {
		sum += index
	}
	return sum
}
