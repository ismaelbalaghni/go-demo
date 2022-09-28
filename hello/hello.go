package main

import (
	"fmt"
	"rsc.io/quote"
	"example/greetings"
)

func main(){
	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	// Utilisation du module greetings
	message := greetings.Hello("Mark")
	fmt.Println(message)
}
