package main

import (
	"fmt"
	"time"
)

const red string = "\033[31m"

func main() {
	var i int = 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is something else")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
	timeNow()
	something()
}

func timeNow() {
	var timeNow time.Time = time.Now()
	switch {
	case timeNow.Hour() < 12:
		fmt.Println("Good morning!")
	case timeNow.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Printf("%sGood evening!\n", red)
	}
}

func something() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(42)
	whatAmI("hello")
}
