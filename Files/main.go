package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in the file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	content, err := os.ReadFile(filename)
	checkNilErr(err)
	// str := string(content)

	fmt.Println("Text data inside the file is \n", string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
