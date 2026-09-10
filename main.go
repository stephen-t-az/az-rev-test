package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	var testConfig map[string]string

	fmt.Println("Setting up configuration...")

	testConfig["env"] = os.Getenv("test")

	println("test the review")

	fmt.Println(testConfig["env"])
}
