package main

import (
	"fmt"
	"math"
)

// This would be considered a global variable?  (I know the scope would be globally)
const myName string = "Andrew"

func main() {

	fmt.Println(myName)

	/**
	  * Looks like you can't use := operator with a constant as well...
	  * const number := 50000
	**/

	const n = 5000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
