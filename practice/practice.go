package main

import (
	// "encoding/json"
	"fmt"
	"strconv"
	// "math"
)

func main1() {
	fmt.Println("This is a practice file\n")
	// fmt.Println("var, const, :=")

	// var name string = "Rakshit"
	// const name2 = -2000

	// var x int64 = math.MaxInt64
	// fmt.Println(name, x)
	// test()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// var names = []int{1, 5}
	// names = append(names, 6)
	// var names2 = [5]int{1, 5}
	// var names3 = make([]int, 8, 10)
	// // for names {

	// // }

	// for i, value := range names {
	// 	fmt.Println(i, value)
	// }
	// for i, value := range names2 {
	// 	fmt.Println(i, value)
	// }
	// for i, value := range names3 {
	// 	fmt.Println(i, value)
	// }

	// type Person struct {
	// 	FirstName string
	// 	LastName  string
	// 	Age       int
	// }

	// var Rakshit Person
	// fmt.Println(Rakshit)
	// Rakshit.FirstName = "Rakshit"
	// Rakshit.LastName = "Singh"
	// fmt.Println(Rakshit)

	// var Rakshit1 = new(Person)
	// fmt.Println(Rakshit1)

	// type Person2 struct {
	// 	FirstName *string `json:"first_name"`
	// 	Salary    *int    `json:"salary"`
	// }

	// first := "Rakshit"
	// var salary *int = nil

	// p := Person2{
	//     FirstName: &first,
	//     Salary:    salary,
	// }

	// var p Person2

	// data, _ := json.Marshal(p)
	// fmt.Println(string(data))
	// fmt.Println(p)
	// num := "Rakshit"
	// var ptr *string = nil
	// var name *string = nil
	// fmt.Println("\nInitial")
	// fmt.Println(name)
	// fmt.Println(num)
	// fmt.Println(&num)
	// fmt.Println(name == nil)
	// ptr = &num
	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	// fmt.Println("\n"+`part 2: num = num + " Singh"`)

	// num = num + " Singh"
	// fmt.Println(name)
	// fmt.Println(num)
	// fmt.Println(&num)
	// fmt.Println(name == nil)
	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	// fmt.Println("\n"+`part 3: *ptr = *ptr + " Singh"`)

	// *ptr = *ptr + " Singh"
	// fmt.Println(name)
	// fmt.Println(num)
	// fmt.Println(&num)
	// fmt.Println(name == nil)
	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	// a := 42
	// b := 3.14
	// c := "Hello"
	// d := true
	// ptr := &a

	// type square struct{
	// 	side int
	// }

	// var s square;
	// // s := square{side: 10}
	// s.side=10

	// fmt.Println("\nT verb for type")
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", c)
	// fmt.Printf("%T\n", d)
	// fmt.Printf("%T\n", ptr)
	// fmt.Printf("%T\n", s)

	// fmt.Println("\nv verb for value")
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)
	// fmt.Printf("%v\n", ptr)
	// fmt.Printf("%v\n", s)

	// fmt.Println("\n+v verb for value")
	// fmt.Printf("%+v\n", a)
	// fmt.Printf("%+v\n", b)
	// fmt.Printf("%+v\n", c)
	// fmt.Printf("%+v\n", d)
	// fmt.Printf("%+v\n", ptr)
	// fmt.Printf("%+v\n", s)

	// fmt.Println("\n#v verb for value")
	// fmt.Printf("%#v\n", a)
	// fmt.Printf("%#v\n", b)
	// fmt.Printf("%#v\n", c)
	// fmt.Printf("%#v\n", d)
	// fmt.Printf("%#v\n", ptr)
	// fmt.Printf("%#v\n", s)

	// fmt.Println("\n"+`%% verb for value`)
	// // fmt.Printf("%%\n", a)
	// // fmt.Printf("%%\n", b)
	// // fmt.Printf("%%\n", c)
	// // fmt.Printf("%%\n", d)
	// // fmt.Printf("%%\n", ptr)
	// // fmt.Printf("%%\n", s)
	// // fmt.Printf("\n%%")
	// // fmt.Printf("%%\n")

	// fmt.Print("Line1")
	// fmt.Printf("Line2")
	// fmt.Printf("Line3")

	var a = "5"
	// var b = float64(a) // wrong
	b, err := strconv.ParseFloat(a, 64)
	if err == nil {
		fmt.Println(b)
	} else {
		fmt.Println(err)
	}

	c, err := strconv.Atoi(a)

	if err == nil {
		fmt.Println(c)
	} else {
		fmt.Println(err)
	}

	d := strconv.Itoa(c)
	fmt.Println(d)

}

// func test() {
// 	truth := true
// 	truth1 := false
// 	fmt.Println(truth, truth1)
// 	truth = false
// 	fmt.Println(truth, truth1)
// }

func main() {
	main1()
}
