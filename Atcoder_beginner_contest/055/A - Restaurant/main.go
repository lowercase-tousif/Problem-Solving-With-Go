package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	x := a * 800
	y := 200
	z := a / 15

	res := 0

	if z == 1 {
		res = x - y
		fmt.Println(res)
	} else {
		y = z * y
		res = x - y
		fmt.Println(res)
	}

}
