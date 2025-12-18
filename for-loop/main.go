package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Numbers is :", i)
	// }


	// // Infinite loof
	// counter := 0
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	counter++
	// 	if counter == 70 {
	// 		break
	// 	}
	// }

	// numbers := []int{11, 42, 83, 14, 75}
	// for index, value := range numbers {
	// 	fmt.Printf("Index of Numbers is %d, and value is %d\n", index, value)
	// }

	data := "Hello, world!"
	for index, char := range data {
		fmt.Printf("Index of Data is %d, and value is %c\n", index, char)
	}

}
