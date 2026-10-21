package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Parse() (baseUrl, apiKey, modelName string, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return "", "", "", err
	}

	modelProfile := os.Getenv("MODEL_PROFILE")
	baseUrl = os.Getenv(fmt.Sprintf("%s_BASE_URL", modelProfile))
	apiKey = os.Getenv(fmt.Sprintf("%s_API_KEY", modelProfile))
	modelName = os.Getenv(fmt.Sprintf("%s_MODEL_NAME", modelProfile))
	return
}
