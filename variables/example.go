package main

import (
	"fmt"
)

func main() {
	/**
	 * Alright so here is the deal about this... I know that you can infer types, but no.
	 * @Override - but here is an example of inferred-type and static-typed.
	**/

	//Inferred. <{^.^}> *"uWu"*
	var a = "something" // Hey this knows I'm a string :)
	fmt.Printf("a:%s\n", a)
	// Static 8)
	var ay string = "String" // I KNOW I AM A STRING
	fmt.Printf("'a: %s\n", ay)

	// You can do multple variables on one line
	var b, c int = 2, 1 // b = 2; c = 1; (ayo)
	fmt.Printf("b: %d\nc: %d\n", b, c)

	var tool bool = true
	fmt.Println(tool) // true

	/**
		* Alright this is the EXCEPTION -> Just because I like this syntax lmao
	  * @Override - only flaw, you can't use types at all with := operator...
	**/

	fruit := "apple" //we know it's a string lmao

	fmt.Printf("Fruit: %s\n", fruit)

	// initializing variables

	var something float64
	fmt.Println(something)

}
