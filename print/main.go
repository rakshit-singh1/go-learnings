package main

import "fmt"

func main() {
	age := 26
	name := "Rakshit"
	height := 5.9234567

	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	// age:  26 height:  5.9234567 name:  Rakshit
	fmt.Println("Hello world")

	// fmt.Printf("age is %d\n", age)
	// fmt.Printf("height is %.2f\n", height)
	// fmt.Printf("Type of age is %T\n", age)
	// fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
