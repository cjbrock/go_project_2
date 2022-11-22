package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	people := []string{"Anna", "Todd", "Corinna"}

	messages, err := greetings.Hellos(people)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
