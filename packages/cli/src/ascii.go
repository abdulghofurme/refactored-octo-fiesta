package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World")
	data, err := os.ReadFile("src/cow.txt")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))
}
