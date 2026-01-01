package main

import (
	"fmt"
	"strings"
)

func main() {
	datam := "a,b,c, d"
	var elements = strings.Split(datam, ",")
	fmt.Println(strings.TrimSpace(elements[len(elements)-1]))

	count1 := strings.Count(datam, ",")
	fmt.Println("count1: ", count1)

	data := "apple.orange.banana.prince"
	parts := strings.Split(data, ".")
	fmt.Println(parts[0])

	str := "one two three four two two five"
	count := strings.Count(str, "tw")
	fmt.Println("count: ", count)

	str = "                       Hello,              Go!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed: ", trimmed)

	str1 := "Rakshit"
	str2 := "Singh"
	result := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("result: ", result)
}
