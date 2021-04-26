package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	var number1 int = 10
	fmt.Println(number1)

	var a, b int
	a = 100
	b = 200

	fmt.Println(a, b)

	var n1, n2 = 300, 400
	fmt.Println(n1, n2)

	var (
		name    string
		address string
		age     int
	)

	name = "vinh"
	address = "Ha noi"
	age = 35

	fmt.Println(name, address, age)

	var name2, address2, age2 = "name2", "address2", 35
	fmt.Println(name2, address2, age2)

}
