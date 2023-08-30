package main

import (
	"context"
	"encoding/json"

    "strconv"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	

)

type User struct {
	Email string `json:"email"`
	Username string `json:"username"`
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
		return events.APIGatewayProxyResponse{Body: "Error parsing request body", StatusCode: 400 ,  Headers: map[string]string{
            "Access-Control-Allow-Origin":"*",
            "Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Methods":     "POST, GET, OPTIONS", // Add any other methods you'd want to support
    		"Access-Control-Allow-Headers":     "Content-Type, Authorization", // Add other headers you might be sending in requests
        },}, err
	}

	// Create user logic
	user, err = CreateUser(ctx, user)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error creating user", StatusCode: 500,  Headers: map[string]string{
            "Access-Control-Allow-Origin":"*",
            "Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Methods":     "POST, GET, OPTIONS", // Add any other methods you'd want to support
    		"Access-Control-Allow-Headers":     "Content-Type, Authorization", // Add other headers you might be sending in requests
        },}, err
	}

	body, err := json.Marshal(user)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error marshalling response", StatusCode: 500,  Headers: map[string]string{
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS", // Add any other methods you'd want to support
    		"Access-Control-Allow-Headers": "Content-Type, Authorization", // Add other headers you might be sending in requests
        },}, err
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200, Headers: map[string]string{
        "Access-Control-Allow-Origin":  "*",
        "Access-Control-Allow-Credentials":"true",
		"Access-Control-Allow-Methods":"POST, GET, OPTIONS", // Add any other methods you'd want to support
    	"Access-Control-Allow-Headers":"Content-Type, Authorization", // Add other headers you might be sending in requests
    }}, nil
}

func CreateUser(ctx context.Context, user User) (User, error) {

	// Defaults for Speed and Pitch if they are not set
	if user.Speed == 0 {
		user.Speed = 50
	}
	if user.Pitch == 0 {
		user.Pitch = 50
	}

	


	input := &dynamodb.PutItemInput{
		TableName: aws.String("OTtable"),
		Item: map[string]types.AttributeValue{
			"Email":    &types.AttributeValueMemberS{Value: user.Email},
			"Username": &types.AttributeValueMemberS{Value: user.Username},
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
