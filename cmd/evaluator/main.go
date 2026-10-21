package main

import (
	"fmt"
	"log"

	cl "github.com/grafviktor/openai-classifier/internal/classifier"
	"github.com/grafviktor/openai-classifier/internal/configuration"
	"github.com/grafviktor/openai-classifier/internal/evaluator"
)

type mismatch struct {
	id       int
	text     string
	expected cl.Classification
	actual   cl.Classification
}

func main() {
	baseUrl, apiKey, modelName, err := configuration.Parse()
	if err != nil {
		log.Fatal(err)
	}

	classifier := cl.NewClassifier(baseUrl, apiKey, modelName)
	evaluationDataSet := evaluator.LoadDataset()
	datasetLength := len(evaluationDataSet)

	totalEvaluated := datasetLength
	correctService := datasetLength
	correctPriority := datasetLength
	correctCategory := datasetLength

	var mismatches []mismatch

	for index, dataItem := range evaluationDataSet {
		// if index > 3 {
		// 	break
		// }

		mismatchFound := false

		expectedClassification := dataItem.Classification
		actualClassification, _ := classifier.Classify(dataItem.Text)

		if expectedClassification.Service != actualClassification.Service {
			correctService--
			mismatchFound = true
		}

		if expectedClassification.Priority != actualClassification.Priority {
			correctPriority--
			mismatchFound = true
		}

		if expectedClassification.Category != actualClassification.Category {
			correctCategory--
			mismatchFound = true
		}

		if mismatchFound {
			mismatches = append(mismatches, mismatch{
				id:       index,
				text:     dataItem.Text,
				expected: expectedClassification,
				actual:   actualClassification,
			})
		}
	}

	fmt.Println("\n-----------------")
	fmt.Printf("Evaluated: %d\n\n", totalEvaluated)
	fmt.Printf("Service:\n  correct: %d (%d%%)\n", correctService, calcPercentage(totalEvaluated, correctService))
	fmt.Printf("Priority:\n  correct: %d (%d%%)\n", correctPriority, calcPercentage(totalEvaluated, correctPriority))
	fmt.Printf("Category:\n  correct: %d (%d%%)\n", correctCategory, calcPercentage(totalEvaluated, correctCategory))
	printMismatches(mismatches)
}

func calcPercentage(total, correct int) int {
	return correct / total * 100
}

func printMismatches(mismatches []mismatch) {
	fmt.Println("\nErrors: ")

	for _, m := range mismatches {
		fmt.Printf("  #%d\n", m.id)
		fmt.Printf("  Input: %s\n", m.text)
		fmt.Printf("  Expected: %s\t%s\t%s\n", m.expected.Service, m.expected.Priority, m.expected.Category)
		fmt.Printf("  Actual:   %s\t%s\t%s\n\n", m.actual.Service, m.actual.Priority, m.actual.Category)
	}
}
