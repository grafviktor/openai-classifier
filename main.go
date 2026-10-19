package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dispatcher := NewDispatcher()
	classifier, err := NewOpenAIClassifier("gpt-6-luna")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			break
		}

		if err := scanner.Err(); err != nil {
			log.Println(err.Error())
		}

		userInput := scanner.Text()
		classification, err := classifier.Classify(userInput)
		if err != nil {
			log.Printf("Error message: %v", err)
			continue
		}
		dispatcher.Dispatch(Message{
			classification,
			userInput,
		})
	}
}
