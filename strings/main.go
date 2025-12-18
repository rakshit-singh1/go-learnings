package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple.orange.banana.Rakshit"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count: ", count)

	str = "                       Hello,              Go!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed: ", trimmed)

	str1 := "Rakshit"
	str2 := "Singh"
	result := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("result: ", result)
}
