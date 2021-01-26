package main

import "fmt"

func  main(){
	// this is the first program by greeting the world
	fmt.Println("Hello World !")

	// testing out back quote which is raw literal
	a := `say "Hello" to Go!`
	b := `Testing this again "I know" I want it on another line "\n"`

	fmt.Println(a)
	fmt.Println(b)
}
