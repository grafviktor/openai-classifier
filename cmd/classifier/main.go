package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	cl "github.com/grafviktor/openai-classifier/internal/classifier"
	"github.com/grafviktor/openai-classifier/internal/configuration"
)

func main() {
	baseUrl, apiKey, modelName, err := configuration.Parse()
	if err != nil {
		log.Fatalf("load .env: %v", err)
	}

	classifier := cl.NewClassifier(baseUrl, apiKey, modelName)
	dispatcher := cl.NewDispatcher()
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
		log.Println("Classification in progress")
		classification, err := classifier.Classify(userInput)
		if err != nil {
			log.Printf("Error message: %v", err)
			continue
		}
		log.Println("Classified")
		dispatcher.Dispatch(cl.Message{
			Classification: classification,
			Text:           userInput,
		})
	}
}
