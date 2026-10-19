## Configuration ##

All models should be defined in .env file in the following format:

```dotenv
# Model profile which will be used when the app starts
MODEL_PROFILE=LOCAL_QWEN34B

# Model profile: LOCAL_QWEN34B
LOCAL_QWEN34B_BASE_URL=http://localhost:11434/v1
LOCAL_QWEN34B_API_KEY=ollama
LOCAL_QWEN34B_MODEL_NAME=qwen3:4b

# Model profile:  REMOTE_GPT5MINI
REMOTE_GPT5MINI_BASE_URL=https://api.openai.com/v1
REMOTE_GPT5MINI_API_KEY=...
REMOTE_GPT5MINI_MODEL_NAME=gpt-5-mini

# Model profile: REMOTE_GPT6LUNA
REMOTE_GPT6LUNA_BASE_URL=https://api.openai.com/v1
REMOTE_GPT6LUNA_API_KEY=...
REMOTE_GPT6LUNA_MODEL_NAME=gpt-6-luna
```

## Reasoning effort ##

That's related to ollama configuration. By default if a model supports thinking, and you can check this with this command:

```bash
 ollama show qwen3:4b

 Model
    architecture        qwen3
    parameters          4.0B
    context length      262144
    embedding length    2560
    quantization        Q4_K_M

  Capabilities
    completion
    tools
    thinking            <-------- Thinking!
```
... ollama will enable this thinking mode by default. As a result, even if your model must only label something, behind the scene it will also reasons its choice, which takes time.

### Option 1 ###

Find a model without thinking feature.
```bash
ollama show qwen3:4b-instruct
  Model
    architecture        qwen3
    parameters          4.0B
    context length      262144
    embedding length    2560
    quantization        Q4_K_M

  Capabilities
    completion
    tools
                        <-------- Not thinking!
```

### Option 2 ###

Probably this will not work

Disable reasoning by adding a similar flag in .env file:
```dotenv
LOCAL_QWEN34B_REASONING_EFFORT=none
```
And then set a correct parameter in OpenAI response params:
```go
	responseParams := responses.ResponseNewParams{...}
	if c.reasoningEffort != "" {
		responseParams.Reasoning = shared.ReasoningParam{
			Effort: shared.ReasoningEffort(c.reasoningEffort),
		}
	}
```

For completeness, these are the model reasoning options:

|**Value**|**Meaning**                        |
|---------|-----------------------------------|
| none    | Do not reason                     |
| minimal | Smallest reasoning effort         |
| low     | Light reasoning                   |
| medium  | Default for most reasoning models |
| high    | More reasoning                    |
| xhigh   | Extra-high reasoning              |
| max     | Maximum reasoning                 |