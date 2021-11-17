package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main()  {

	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func (bot englishBot) getGreeting() string {
	return "Hi There!"
}

func (bot spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(bot bot)  {
	fmt.Println(bot.getGreeting())
}
