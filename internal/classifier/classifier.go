package classifier

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

var _ Classifier = (*OpenAIClassifier)(nil)

type Classifier interface {
	Classify(text string) (Classification, error)
}

const systemInstruction = `
Classify the user's message based only on its meaning.

Service values:
- payment: problems with payments, cards, transactions, transfers
- authentication: login, password, MFA, account access
- storage: files, storage capacity, data storage
- unknown: only if you cannot determine the service. Most likely user input error.

Priority values:
- low: informational request or issue with little immediate impact. Use it if service or category is unknown.
- medium: user is affected but there is a workaround or limited impact
- high: critical functionality is unavailable, a transaction is blocked, or there is significant business impact

Category values:
- incident: something is broken or not working as expected
- request: the user wants something to be done or changed
- question: the user is asking for information
- unknown: only if you cannot categorize this input. Most likely user input error.

Do not infer information that is not reasonably supported by the user's message.
If several classifications are possible, choose the one that best matches the primary intent.
Return only the requested structured classification.
`

type OpenAIClassifier struct {
	client *openai.Client
	model  string
}

func NewClassifier(baseUrl, apiKey, modelName string) OpenAIClassifier {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseUrl),
	)

	return OpenAIClassifier{
		client: &client,
		model:  modelName,
	}
}

func (c OpenAIClassifier) Classify(text string) (Classification, error) {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	schema := reflector.Reflect(new(Classification))
	schemaData, err := json.Marshal(schema)
	if err != nil {
		return Classification{}, fmt.Errorf("build schema: %w", err)
	}

	var schemaMap map[string]any
	if err := json.Unmarshal(schemaData, &schemaMap); err != nil {
		return Classification{}, fmt.Errorf("decode schema: %w", err)
	}

	input := responses.ResponseNewParamsInputUnion{
		OfInputItemList: responses.ResponseInputParam{
			responses.ResponseInputItemParamOfMessage(
				systemInstruction,
				// Role developer is only support by open AI models, it's safer to stick to role 'System'
				// responses.EasyInputMessageRoleDeveloper,
				responses.EasyInputMessageRoleSystem,
			),
			responses.ResponseInputItemParamOfMessage(
				text,
				responses.EasyInputMessageRoleUser,
			),
		},
	}

	responseParams := responses.ResponseNewParams{
		Model: c.model,
		Input: input,
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigUnionParam{
				OfJSONSchema: &responses.ResponseFormatTextJSONSchemaConfigParam{
					Name:        "classification",
					Description: openai.String("Classify a user request"),
					Schema:      schemaMap,
					Strict:      openai.Bool(true),
				},
			},
		},
	}

	resp, err := c.client.Responses.New(context.Background(), responseParams)
	if err != nil {
		return Classification{}, fmt.Errorf("response error: %w", err)
	}

	var message Classification

	err = json.Unmarshal([]byte(resp.OutputText()), &message)
	if err != nil {
		return Classification{}, fmt.Errorf("decode classification: %w", err)
	}

	if err := message.Validate(); err != nil {
		return Classification{}, fmt.Errorf(
			"invalid classification: %w", err,
		)
	}

	return message, nil
}
