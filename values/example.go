package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello" + " World")
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	fmt.Printf("7.0 / 3.0 = %.2f\n", 4.0/2.0)
	fmt.Println(!true) // false
}
