package main

import (
	"fmt"
	"math"
)

const red string = "\033[31m"

var something int32

func main() {
	logic()
}

func logic() {
	recursive := recursiveDecrement(55)
	fmt.Println(recursive)
	fmt.Println("Testing if the number is even: ", testIfIsEven(5))
	inputHandler()

}
func inputHandler() {
	fmt.Println("Please enter a number: ")
	fmt.Scan(&something)
	anotherFunctionComparisonButWithAnd(something)
}

func testIfIsEven(n int32) int32 {
	if n%2 == 0 {
		fmt.Println("\032[33m passed!\033[0m")
		return n
	}
	return 0
}

func recursiveDecrement(n int32) int32 {
	if n < 1 {
		fmt.Println("\033[34mYo we actually worked\033[0m")
		return 0
	}
	fmt.Printf("We are going through recursion\nn = %d\n", n)
	return recursiveDecrement(n - 1)
}

// this is showing you a conditional with an or
func hereIsAnotherTest(number int32) {
	if number%2 == 0 || number%10 == 0 {
		fmt.Println("\033[31mSuccess!\n\033[0mThe nunmber is divisible by either 2 or 10")
	} else {
		fmt.Println("sorry you don't have the right number for this.")
	}

}

// this is showing you a conditional with an andor
func anotherFunctionComparisonButWithAnd(number int32) {
	if number == 15 {
		value := math.Sin(float64(number))
		fmt.Println(value)
	} else if number == 20 {
		var valueTwo float64 = math.Sin(float64(number))
		fmt.Println(valueTwo)
	} else if number > 15 {
		fmt.Println("\033[34mGreat Success!")
	} else {
		fmt.Printf("%sWe didn't print out shit my boi lmao", red)
	}
}
