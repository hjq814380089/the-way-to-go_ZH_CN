package main

import (
	"strings"
	"fmt"
)

	func main() {
	s := "Hello, 世界! Hello!"
	ss := strings.Fields(s)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
}

