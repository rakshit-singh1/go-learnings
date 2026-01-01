package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	// myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	myURL := "wss://example.com:8080/path/to/resource" // Scheme is necessary
	fmt.Printf("Type of URL: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of parsed URL: %T\n", parsedURL)
	fmt.Println("Parsed URL:", parsedURL)
	// fmt.Println()
	// fmt.Printf("Type of Scheme: %T\n", parsedURL.Scheme)
	// fmt.Printf("Type of Host: %T\n", parsedURL.Host)
	// fmt.Printf("Type of Path: %T\n", parsedURL.Path)
	// fmt.Printf("Type of RawQuery: %T\n", parsedURL.RawQuery)
	fmt.Println()
	fmt.Println("Scheme", parsedURL.Scheme)
	fmt.Println("Host", parsedURL.Host)
	fmt.Println("Path", parsedURL.Path)
	fmt.Println("RawQuery", parsedURL.RawQuery)

	// parsedURL, err := url.Parse(myURL)
	// if err != nil {
	// 	fmt.Println("Can't parse URL ", err)
	// 	return
	// }
	// // fmt.Printf("Type of URL: %T\n", parsedURL)

	// fmt.Println("Scheme: ", parsedURL.Scheme)
	// fmt.Println("Host: ", parsedURL.Host)
	// fmt.Println("Path: ", parsedURL.Path)
	// fmt.Println("RawQuery: ", parsedURL.RawQuery)

	// Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iamRakshit"
	fmt.Println("Parsed URL:", parsedURL)
	// Constructing a URL string from a URL object
	newUrl := parsedURL.String()

	fmt.Println("new URL: ", newUrl)

}
