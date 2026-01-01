package main

import (
	// "encoding/json"
	"fmt"
	"io"
	// "reflect"
	// "strconv"
	// "strings"

	// "io/ioutil" // depricated
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // response is a struct
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// fmt.Printf("%T\n", response)
	// fmt.Println(response)




	// fmt.Println("\nExpanding response1")
	// fmt.Println("Status Code:", response.StatusCode) // 200
	// fmt.Println("Status:", response.Status)          // "200 OK"
	// fmt.Println("Proto:", response.Proto)            // HTTP/1.1
	// for key, values := range response.Header {
	// 	for _, value := range values {
	// 		fmt.Printf("%s: %s\n", key, value)
	// 	}
	// }
	// statusText := strings.TrimPrefix(response.Status, strconv.Itoa(response.StatusCode)+" ")
	// fmt.Println("statusText",statusText)





	// fmt.Println("\nExpanding response2")
	// v := reflect.ValueOf(*response)
	// t := v.Type()

	// for i := 0; i < v.NumField(); i++ {
	// 	fieldName := t.Field(i).Name
	// 	fieldValue := v.Field(i).Interface()
	// 	fmt.Printf("%s: %v\n", fieldName, fieldValue)
	// }

	// var datam any
	// if err := json.NewDecoder(response.Body).Decode(&datam); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(datam)
	// fmt.Printf("Body type: %T\n", datam)








	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nResponse Body",string(body))

	// var data any
	// if err := json.Unmarshal(body, &data); err != nil { // json.Unmarshal takes pointer type variable in target
	// 	panic(err)
	// }
	// pretty, _ := json.MarshalIndent(data, "", "  ")
	// fmt.Println(string(pretty))

	// res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err != nil {
	// 	fmt.Println("Error getting GET response ", err)
	// 	return
	// }
	// defer res.Body.Close()
	// fmt.Printf("Type of response: %T\n", res)
	// // fmt.Println("response: ", res)

	// // Read the response body
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response ", err)
	// 	return
	// }

	// fmt.Println("response: ", string(data))

}
