package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// differ is printed in reverse order basically follow stack data structure
func main() {

	// // fmt.Println("Start")
	// // defer fmt.Println("In process")
	// // fmt.Println("End")

	// fmt.Println("Starting of the program")
	// data := add(5, 6)
	// defer fmt.Println("Data is : ", data)
	// defer fmt.Println("Middle of the program")
	// fmt.Println("End of the program")
	// // Start
	// // End
	// // Starting of the program
	// // End of the program
	// // Middle of the program
	// // Data is :  11
	// // In process

	// x := 10
	// defer fmt.Println(x)
	// x = 20
	// // 10

	// for i := 0; i < 9; i++ {
	// 	defer fmt.Print(i)
	// }
	// // 876543210

	fmt.Println("Start")
	panic("something went wrong")
	fmt.Println("End")

	// Start
	// panic: something went wrong

	// goroutine 1 [running]:
	// main.main()
	//         /home/blockchain-dev/Downloads/gol_lang_basics/defer/main.go:39 +0x59
	// exit status 2

	fmt.Println("Start")
	defer panic("something went wrong")
	fmt.Println("End")

	// Start
	// panic: something went wrong

	// goroutine 1 [running]:
	// main.main()
	//         /home/blockchain-dev/Downloads/gol_lang_basics/defer/main.go:40 +0x59
	// exit status 2

}
