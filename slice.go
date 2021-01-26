package main

import "fmt"

func  main(){

	b := []int{-1,3, 4, 3, 7}
	fmt.Println(b)

	seaCreatures := []string{"Shark", "cuttlefish", "Squid", "mantis shrimp"}
	
	seaCreatures = append(seaCreatures, "seahorse")
	fmt.Println(seaCreatures)
}
