package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env: %w", err)
	}

	modelProfile := os.Getenv("MODEL_PROFILE")
	baseUrl := os.Getenv(fmt.Sprintf("%s_BASE_URL", modelProfile))
	apiKey := os.Getenv(fmt.Sprintf("%s_API_KEY", modelProfile))
	modelName := os.Getenv(fmt.Sprintf("%s_MODEL_NAME", modelProfile))

	classifier := NewClassifier(baseUrl, apiKey, modelName)
	dispatcher := NewDispatcher()
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
		dispatcher.Dispatch(Message{
			classification,
			userInput,
		})
	}
}
