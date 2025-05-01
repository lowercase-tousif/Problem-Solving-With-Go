package main

import "fmt"

func solve() {
	var c1 int64
	var c2 int64

	var n int64
	fmt.Scan(&n)
	
	var div int64 = n / 3

	c1 = div
	c2 = c1

	if n%3 == 1 {
		c1 += 1
	} else if n%3 == 2 {
		c2 += 1
	}

	fmt.Println(c1, c2)
}

func main() {
	var t int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		solve()

	}
}
