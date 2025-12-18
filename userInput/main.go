package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's your name?")

	// Method 1: Reads only till the first space

	// var name string;
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name);

	// Method 2: Reads the whole line including spaces

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // Reads input till the newline character
	fmt.Println("Hello, Mr.", name)
}
