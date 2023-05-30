package chatgpt

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

type ChatgptClient struct {
	*openai.Client
}

func (c *ChatgptClient) GenerateResponse(prompt string) string {
	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	response, err := c.CreateChatCompletion(context.Background(), request)

	if err != nil {
		log.Println(err.Error())
		return "Espera pe causa mi servidor cuesta :>"
	}

	responseContent := response.Choices[0].Message.Content
	return responseContent
}

func NewClient(authToken string) ChatgptClient {
	defer log.Println("Chatgpt client created successfully")

	return ChatgptClient{
		openai.NewClient(authToken),
	}
}
