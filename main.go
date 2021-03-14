package main

import "fmt"

// all types with the declared function name with the same return value,
// are honorary members of type bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// we want this function to work with both structs. time for interfaces
/*func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}*/

// fixed!
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// since we  are not using the reciever variable iside functions we can omit them in the reciever statement
func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
