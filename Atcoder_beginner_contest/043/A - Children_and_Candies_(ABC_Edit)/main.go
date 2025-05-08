package main

import "fmt"

func main(){
	var a int 
	var sum  = 0
	fmt.Scan(&a)

	for i := 1; i <= a; i++{
		sum += i
	}

	fmt.Println(sum)
}
