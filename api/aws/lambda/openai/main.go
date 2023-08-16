package main

import (
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
	"os"
)

var openaiClient *openai.Client;

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	openaiClient = openai.NewClient(apiKey)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var text = request.Body

	response, err := generateResponse(text)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: response, StatusCode: 200}, nil
}

func generateResponse(text string) (string, error) {

	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func main() {
	lambda.Start(Handler)
}