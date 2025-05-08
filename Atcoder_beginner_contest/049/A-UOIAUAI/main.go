package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string

	fmt.Scan(&a)

	a = strings.ToLower(a)

	if a == "a" || a == "e" || a == "i" || a == "o" || a == "u" {
		fmt.Println("vowel")
	} else {
		fmt.Println("consonant")
	}
}
