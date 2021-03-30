package main

import "fmt"

/**
Interface Bot
EnglishBot implements Bot
SpanishBot implements Bot

Both implements internally because both has all the functions defined in Bot interface
 */

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct {}

func (eb englishBot) getGreeting() string  {
	return "Hello !!"
}

func (sb spanishBot) getGreeting() string  {
	return "Hola !!"
}

func printGreeting(b bot)  {

	// call greeting functions from struct objects which implements bot interface
	fmt.Println(b.getGreeting())
}

func implementBot(){
	fmt.Println("--------------- BOT -----------------")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}