package main

//Remember though if you wanted to space this out, you can name the packages whatever... but MAIN is the MAIN one

import (
	"fmt"
)

func main() {

	/**
		  * Okay so this is where I really lock in, because this is a little different from java
	    * Obv, we know java has { for (int i = 0; i < something; i++) - pre condition}
	    * we also have: {for (something : somethings) - pre condition} (this is the {for-each blank in blank})
		  *
		**/

	i := 0

	for i <= 4 {
		fmt.Println("hello ", i)
		i = i + 1
	}

	// This is what Java heads love and know without () obviously but it's the same semantically.

	for j := 0; j < 4; j++ {
		fmt.Printf("Hello, j: %d\n", j)
	}

	//boom just found out this is not inclusive :p -> [0->3] ![0->4]
	for i := range 4 {
		if i == 4 {
			fmt.Printf("Done with i: %d, leaving loop\n\n", i)
		}
		fmt.Printf("Range: %d\n", i)
	}
	for {
		fmt.Println("Loooping")
		break //need to use the break keyword to terminate the loop :o - this is pretty cool
	}

	//This is going to look odly similar to Python :)
	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("n: %d\n", n)
	}

}
