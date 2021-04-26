package main

import "fmt"

func main() {
	for number := 1; number < 100; number++ {
		if number%15 == 0 {
			fmt.Println("FIZZBUZZ")
		} else if number%3 == 0 {
			fmt.Println("FIZZ")
		} else if number%5 == 0 {
			fmt.Println("BUZZ")
		} else {
			fmt.Println(number)
		}
	}
}
