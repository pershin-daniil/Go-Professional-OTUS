package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const helloMessage = "Hello, OTUS!"

func main() {
	revMessage := stringutil.Reverse(helloMessage)
	fmt.Printf("%s", revMessage)
}
