package main

import (
	"fmt"
	"time"
)

func fizzBuzzSwitch(number int) {
	switch {
	case number%15 == 0:
		fmt.Println("FIZZBUZZ")
	case number%5 == 0:
		fmt.Println("BUZZ")
	case number%3 == 0:
		fmt.Println("FIZZ")
	default:
		fmt.Println(number)
	}
}

func fizzBuzz() {
	for number := 1; number < 100; number++ {
		fizzBuzzSwitch(number)
	}
}

func sayHelloWithTime() {
	var hour int = time.Now().Hour()
	fmt.Println("This is", hour, "hour")
	switch {
	case hour <= 12:
		fmt.Println("Good morning")
	case hour <= 18:
		fmt.Println("Good afternoon")
	case hour < 24:
		fmt.Println("Good evening")
	}
}

func checkType(value interface{}) {
	switch myType := value.(type) {
	case bool:
		fmt.Println(value, "is bool value")
	case int:
		fmt.Println(value, "is integer value")
	case string:
		fmt.Println(value, "is string value")
	default:
		fmt.Println(value, "is unknow datatype", myType)
	}
}

func main() {
	fizzBuzz()
	sayHelloWithTime()
	checkType(true)
	checkType(100)
	checkType("hihi")
	checkType(100.0001)
}
