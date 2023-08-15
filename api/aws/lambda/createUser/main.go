package main

import (
	"context"
	"encoding/json"
	"errors"
    "strconv"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	
	"github.com/google/uuid"
)

type User struct {
	UserId string `json:"userId"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Speed int `json:"speed"`
	Pitch int `json:"pitch"`
}

// Assuming you have initialized this outside in the global scope
var db *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	db = dynamodb.NewFromConfig(cfg)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse body
	var user User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error parsing request body", StatusCode: 400}, err
	}

	// Create user logic
	user, err = CreateUser(ctx, user)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error creating user", StatusCode: 500}, err
	}

	body, err := json.Marshal(user)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error marshalling response", StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200, Headers: map[string]string{
        "Access-Control-Allow-Origin":      "*",
        "Access-Control-Allow-Credentials": "true",
    }}, nil
}

func CreateUser(ctx context.Context, user User) (User, error) {
	if  user.Username == "" || user.Password == "" {
		return User{}, errors.New("Username, and Password are required")
	}

	// Defaults for Speed and Pitch if they are not set
	if user.Speed == 0 {
		user.Speed = 50
	}
	if user.Pitch == 0 {
		user.Pitch = 50
	}

	// Generate a UUID for userId
    userId := uuid.New().String()


	input := &dynamodb.PutItemInput{
		TableName: aws.String("OvertoneTable"),
		Item: map[string]types.AttributeValue{
			"UserId":   &types.AttributeValueMemberS{Value: userId},
			"Email":    &types.AttributeValueMemberS{Value: user.Email},
			"Username": &types.AttributeValueMemberS{Value: user.Username},
			"Password": &types.AttributeValueMemberS{Value: user.Password},
			"Speed":    &types.AttributeValueMemberN{Value: strconv.Itoa(user.Speed)},
			"Pitch":    &types.AttributeValueMemberN{Value: strconv.Itoa(user.Pitch)},
		},
	}	

	_, err := db.PutItem(ctx, input)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func main() {
	lambda.Start(handleRequest)
}
