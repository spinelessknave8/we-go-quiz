package main

import (
	"GO-PROJ/questions"
	"fmt"
)

func main() {
	fmt.Println("This is my quiz game! Let's start playing!")

	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Welcome to the quiz game, %v!\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)
	if age < 10 {
		fmt.Println("Sorry, you're too young to play this game. Please come back when you are older.")
		return
	} else {
		fmt.Printf("Great! You are old enough to play the game, %v.\n", name)
		questions.AskQuestions()
		return
	}

}
