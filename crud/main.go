package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "strings"
)

// node:201 Created is only when you create new record
// json mapped data type
type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	/********************http get request start********************/

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/82")
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 { // res.StatusCode != http.StatusOK
		fmt.Println("Error in getting Response: ", res.Status)
		return
	}

	/********************WE CANNOT READ BODY TWICE********************/

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading : ", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todo Todo

	json.NewDecoder(res.Body).Decode(&todo)
	fmt.Println("Todo", todo)

	/********************http get request end********************/
}
func performPostRequest() {
	/********************http post request start********************/

	todo := Todo{
		UserID:    81,
		Id:        82,
		Title:     "New Task",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	// fmt.Println(jsonData)
	jsonString := string(jsonData)
	// fmt.Println(jsonString)
	jsonReader := strings.NewReader(jsonString)
	// fmt.Println(jsonReader)
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated { // res.StatusCode != 201
		fmt.Println("Error in Posting Response: ", res.Status)
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading create response : ", err)
		return
	}
	fmt.Println("Response: ", string(data))

	/********************http post request start********************/
}
func performPutRequest() {
	/********************http put request start********************/
	todo := Todo{
		UserID:    1,
		Id:        2,
		Title:     "New Task",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/todos/1", jsonReader)

	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 { // res.StatusCode != http.StatusOK
		fmt.Println("Error in Updating Response: ", res.Status)
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading update response : ", err)
		return
	}
	fmt.Println("Response: ", string(data))
	/********************http put request start********************/
}
func performDeleteRequest() {
	req, err := http.NewRequest(
		http.MethodDelete,
		"https://jsonplaceholder.typicode.com/todos/91",
		nil,
	)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Accept all success responses
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		fmt.Println("DELETE failed:", res.Status)
		return
	}

	// 204 has no body → avoid reading
	if res.StatusCode == http.StatusNoContent {
		fmt.Println("Deleted successfully (no content)")
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", string(body))
}
func main() {
	performGetRequest()
	performPostRequest()
	performPutRequest()
	performDeleteRequest()
}

// func performGetRequest() {
// 	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
// 	if err != nil {
// 		fmt.Println("Error getting : ", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		fmt.Println("Error in getting Response: ", res.Status)
// 		return
// 	}

// 	// data, err := ioutil.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	fmt.Println("Error reading : ", err)
// 	// 	return
// 	// }
// 	// fmt.Println("Data: ", string(data))

// 	var todo Todo
// 	err = json.NewDecoder(res.Body).Decode(&todo)
// 	if err != nil {
// 		fmt.Println("Error decoding : ", err)
// 		return
// 	}
// 	fmt.Println("Todo: ", todo)

// 	fmt.Println("Title response : ", todo.Title)
// 	fmt.Println("completed response : ", todo.Completed)
// }

// func performPostRequest() {
// 	todo := Todo{
// 		UserID:    23,
// 		Title:     "Rakshit Kumar",
// 		Completed: true,
// 	}

// 	// Convert the Todo struct to JSON
// 	jsonData, err := json.Marshal(todo)
// 	if err != nil {
// 		fmt.Println("Error marshalling : ", err)
// 		return
// 	}

// 	// convert json data to string
// 	jsonString := string(jsonData)

// 	// convert string data to reader
// 	jsonReader := strings.NewReader(jsonString)

// 	myURL := "https://jsonplaceholder.typicode.com/todos"

// 	// send POST request
// 	res, err := http.Post(myURL, "application/json", jsonReader)
// 	if err != nil {
// 		fmt.Println("Error sending request : ", err)
// 		return
// 	}

// 	defer res.Body.Close()

// 	// data, _ := ioutil.ReadAll(res.Body)
// 	// fmt.Println("Response : ", string(data))

// 	fmt.Println("Response  status : ", res.Status)
// }

// func performUpdateRequest() {
// 	todo := Todo{
// 		UserID:    100001,
// 		Title:     "Rakshit Kumar learning Golang",
// 		Completed: false,
// 	}

// 	// Convert the Todo struct to JSON
// 	jsonData, err := json.Marshal(todo)
// 	if err != nil {
// 		fmt.Println("Error marshalling : ", err)
// 		return
// 	}

// 	// convert json data to string
// 	jsonString := string(jsonData)

// 	// convert string data to reader
// 	jsonReader := strings.NewReader(jsonString)

// 	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

// 	// create PUT Request

// 	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
// 	if err != nil {
// 		fmt.Println("Error creating PUT Request : ", err)
// 		return
// 	}
// 	req.Header.Set("Content-type", "application/json")

// 	// Send the request
// 	client := http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request : ", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	data, _ := ioutil.ReadAll(res.Body)
// 	fmt.Println("Response : ", string(data))
// 	fmt.Println("Response  status : ", res.Status)
// }

// func performDeleteRequest() {
// 	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

// 	// create DELETE Request

// 	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
// 	if err != nil {
// 		fmt.Println("Error creating Delete Request : ", err)
// 		return
// 	}

// 	// Send the request
// 	client := http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request : ", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	fmt.Println("Response  status : ", res.Status)
// }

// func main() {
// 	fmt.Println("Learning CRUD...")
// 	// performGetRequest()
// 	// performPostRequest()
// 	// performUpdateRequest()
// 	performDeleteRequest()
// }
