package main

import "fmt"

func main() {

	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a int
		fmt.Scan(&a)

		fmt.Println(a * 2)
	}

}
