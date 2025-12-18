package main

import "fmt"

func simpleFunction() {
	fmt.Println("My First Function")
}

func add(a int, b int) (result int) {
	result = a + b
	return
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func subtract(a, b int) int {
	return a - b
}

func divide(a, b int) int {
	return a / b
}

func main() {
	simpleFunction()

	addResult := add(4, 2)
	fmt.Println("add of two numbers is:", addResult)

	mul := multiply(4, 2)
	fmt.Println("multiplication of two numbers is:", mul)

	sub := subtract(4, 2)
	fmt.Println("subtraction of two numbers is:", sub)

	div := divide(4, 2)
	fmt.Println("division of two numbers is:", div)
}
