package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofiber/fiber/v2"
)

func main() {

	type User struct {
		UserId string `json:"userId"`
		Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		Speed int `json:"speed"`
		Pitch int `json:"pitch"`
	}

	// Create a new DynamoDB client
	client := startDB()
	
	// Check if the client is nil
	if client == nil {
		// If the client is nil, then the database is not running
		// Print an error message and exit
		fmt.Println("Error: DynamoDB is not running")
		return
	}
	
	// Create a new Fiber app
	app := fiber.New()

	// Create a GET route on path "/api/database"
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		// Get the id from the URL
		id := c.Params("id")
		
		// Create a new GetItemInput
		params := &dynamodb.GetItemInput{
			TableName: aws.String("Users"),
			Key: map[string]*dynamodb.AttributeValue{
				"userId": {
					N: aws.String(id),
				},
			},
		}
		
		// Get the item from DynamoDB
		result, err := client.GetItem(params)
		
		// Check for errors
		if err != nil {
			// If there is an error, print it and return
			fmt.Println(err.Error())
			return err
		}
		
		// Create a new User struct
		user := new(User)
		
		// Unmarshal the result into the User struct
		err = dynamodbattribute.UnmarshalMap(result.Item, user)
		
		// Check for errors
		if err != nil {
			// If there is an error, print it and return
			fmt.Println(err.Error())
			return err
		}
		
		// Return the user as JSON
		return c.JSON(user)
	})

	app.Listen(":3000")
}