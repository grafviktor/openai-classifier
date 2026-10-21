package evaluator

import (
	"encoding/json"
	"os"

	"github.com/grafviktor/openai-classifier/internal/classifier"
)

func LoadDataset() []classifier.Message {
	data, err := os.ReadFile("internal/evaluator/dataset.json")
	if err != nil {
		panic(err)
	}

	var messages []classifier.Message
	if err := json.Unmarshal(data, &messages); err != nil {
		panic(err)
	}
	return messages
}
