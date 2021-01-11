package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGretting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGretting())
}

func (englishBot) getGretting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGretting() string {
	return "Hola!"
}
