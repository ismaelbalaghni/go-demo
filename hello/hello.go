package main

import (
	"bufio"
	"example/greetings"
	"fmt"
	"os"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	// Utilisation du module greetings
	message := greetings.Hello("Mark")
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	nom := scanner.Text()
	if nom == "Joshua" {
		message = greetings.Hello(nom + " Bledsoe")
	} else if nom == "Drew" {
		message = greetings.Hello(nom + " Bledsoe")
	} else {
		fmt.Println("The name you entered is unknown.")
	}
	//message = greetings.Hello(nom)
	fmt.Println(message)
}
