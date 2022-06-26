package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type hindiBot struct{}

func main() {
	eb := englishBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(hb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi!"
}

func (hindiBot) getGreeting() string {
	return "Namaste!"
}