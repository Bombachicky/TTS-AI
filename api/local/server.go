package main

import (
	"context"
	"encoding/json"
	"fmt"

	//"os"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"

	//"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofiber/fiber/v2"
	//"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sashabaranov/go-openai"
)

func parseAPIKey(secretString string) (string, error) {

	size := len(secretString)

	for i := 0; i < size; i++ {
		if secretString[i] == ':' {
			return secretString[i+2:size-2], nil
		}
	}

	return "", nil
}

func getAPIKey() (string, error) {
	
	config, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("Error loading config")
		return "", err
	}

	client := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("OpenAI"),
	}

	secretValue, err := client.GetSecretValue(context.Background(), input)

	if err != nil {
		fmt.Println("Error getting secret value")
		return "", err
	}

	apiKey := *secretValue.SecretString

	apiKey, err = parseAPIKey(apiKey)

	if err != nil {
		fmt.Println("Error parsing API key")
		return "", err
	}

	return apiKey, nil
}

func main() {

	// type User struct {
	// 	UserId string `json:"userId"`
	// 	Email string `json:"email"`
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// 	Speed int `json:"speed"`
	// 	Pitch int `json:"pitch"`
	// }

	// // Create a new DynamoDB client
	// client := startDB()
	
	// // Check if the client is nil
	// if client == nil {
	// 	// If the client is nil, then the database is not running
	// 	// Print an error message and exit
	// 	fmt.Println("Error: DynamoDB is not running")
	// 	return
	// }
	
	// Create a new Fiber app
	app := fiber.New()

	app.Use(cors.New())

	// Create POST request to receive ChatGPT response from user input
	app.Post("/message", func(c *fiber.Ctx) error {

		apiKey, err := getAPIKey()

		if err != nil {
			return c.Status(500).SendString("Error getting API key")
		}

		openaiClient := openai.NewClient(apiKey)

		var input string
		err = json.Unmarshal([]byte(c.Body()), &input)

		if err != nil {
			return c.Status(500).SendString("Error unmarshalling JSON")
		}

		response, err := openaiClient.CreateChatCompletion(
			c.Context(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: input,
					},
				},
			},
		)

		if err != nil {
			return c.Status(500).SendString("Error generating response")
		}

		return c.Status(200).SendString(response.Choices[0].Message.Content)

	})

	app.Listen(":3001")
}