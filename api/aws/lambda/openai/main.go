package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
	openai "github.com/sashabaranov/go-openai"
)

var openaiClient *openai.Client;

func parseAPIKey(secretString string) (string, error) {

	size := len(secretString)

	for i := 0; i < size; i++ {
		if secretString[i] == ':' {
			return secretString[i+2:size-2], nil
		}
	}

	return "", nil
}

func init() {
	config, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("Error loading config")
	}

	client := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("OpenAI"),
	}

	secretValue, err := client.GetSecretValue(context.Background(), input)

	if err != nil {
		fmt.Println("Error getting secret value")
	}

	apiKey := *secretValue.SecretString

	apiKey, err = parseAPIKey(apiKey)

	if err != nil {
		fmt.Println("Error parsing API key")
	}

	openaiClient = openai.NewClient(apiKey)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	text := request.Body
	response, err := generateResponse(text)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error generating response", StatusCode: 500, Headers: map[string]string{
			"Access-Control-Allow-Origin":"*",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS", // Add any other methods you'd want to support
			"Access-Control-Allow-Headers": "Content-Type, Authorization", // Add other headers you might be sending in requests
		}}, err
	}

	return events.APIGatewayProxyResponse{Body: response, StatusCode: 200, Headers: map[string]string{
        "Access-Control-Allow-Origin":  "*",
        "Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Methods": "POST, GET, OPTIONS", // Add any other methods you'd want to support
    	"Access-Control-Allow-Headers": "Content-Type, Authorization", // Add other headers you might be sending in requests
    }}, nil
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