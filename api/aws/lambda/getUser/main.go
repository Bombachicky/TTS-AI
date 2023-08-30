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




func GetUser(ctx context.Context, Username string) (User, error) {
    input := &dynamodb.GetItemInput{
        TableName: aws.String("OTtable"),
        Key: map[string]types.AttributeValue{
            "username": &types.AttributeValueMemberS{Value: Username},
        },
    }

    result, err := db.GetItem(ctx, input)
    if err != nil {
        return User{}, err
    }

    if result.Item == nil {
        return User{}, errors.New("User not found")
    }

    user := User{
    Email: result.Item["Email"].(*types.AttributeValueMemberS).Value,
	}

	speed, err := strconv.Atoi(result.Item["Speed"].(*types.AttributeValueMemberN).Value)
	if err != nil {
		return User{}, err
	}
	user.Speed = speed

	pitch, err := strconv.Atoi(result.Item["Pitch"].(*types.AttributeValueMemberN).Value)
	if err != nil {
		return User{}, err
	}
	user.Pitch = pitch

    return user, nil
}

func handleGetUserRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
   userName, ok := request.PathParameters["username"]
    if !ok || userName == "" {
        return events.APIGatewayProxyResponse{Body: "Username parameter is required", StatusCode: 400}, nil
    }
    user, err := GetUser(ctx, userName)
    if err != nil {
        if err.Error() == "User not found" {
            return events.APIGatewayProxyResponse{Body: "User not found", StatusCode: 404}, nil
        }
        return events.APIGatewayProxyResponse{Body: "Error fetching user", StatusCode: 500}, err
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

//Adjust the main function accordingly if you want to include both handlers
func main() {
    
    lambda.Start(handleGetUserRequest) // For GetUser
}
