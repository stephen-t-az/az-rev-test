package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	test := "First test"

	fmt.Println(test)

	env := os.Getenv("test")

	fmt.Println(env)
}
