package main

import "fmt"

func main() {
		n := 0
		a := n + 1
		
		fmt.Println(a)

		var pointer *int = &a

		fmt.Println(pointer)

}