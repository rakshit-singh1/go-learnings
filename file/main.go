package main

import (
	"fmt"
	"io"
	// "io"
	"os"
)

func AppendToFile(filename, content string) error {
	// Open file in append mode, write-only, create if not exists
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // ensures flush

	// Write content
	_, err = file.WriteString(content)
	return err
}

func main() {
	/**********Open and write***********/
	// file, err := os.Create("example.txt")

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	// fmt.Println(file) // &{0xc000036120}
	// 	content := "hi\nOye balle balle oye"

	// 	byte, errors := io.WriteString(file, content)
	// 	fmt.Println("byte written: ", byte)
	// 	if errors != nil {
	// 		fmt.Println("Error while writing file: ", errors)
	// 		return
	// 	}

	// }

	// defer file.Close()

	/**********Append***********/

	// content := "hi\nOye balle balle oye\n"

	// err := AppendToFile("data.txt", content)
	// if err != nil {
	// 	fmt.Println("Error appending to file:", err)
	// 	return
	// }

	// fmt.Println("Content appended successfully!")

	/**********Reading the file without inbuilt function***********/

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
		}

		// processing file
		// fmt.Print(n) //23

		fmt.Println(buffer[:n])
		fmt.Println(string(buffer[:n]))
	}

	/**********Reading the file with inbuilt function***********/  // Not to be used when size of file is very large

	// content, err := os.ReadFile("example.txt") // previously ioutil.ReadFile was used
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// fmt.Println("Content:",string(content))
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file: ", err)
			return
		}
		defer file.Close()

		content := "Sample Content to write into file"
		byte, errors := io.WriteString(file, content+"\n")
		fmt.Println("byte written: ", byte)
		if errors != nil {
			fmt.Println("Error while writing file: ", errors)
			return
		}

		fmt.Println("sucessfully created file")

	*/

	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file: ", err)
			return
		}
		defer file.Close()

		// Create a buffer to read the file content
		buffer := make([]byte, 1024)

		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file", err)
				return
			}

			// Process the read content
			fmt.Println(string(buffer[:n]))
		}
	*/

	// Read the entire file into a byte slice
	// content, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while reading file ", err)
	// 	return
	// }
	// fmt.Println(string(content))
}
