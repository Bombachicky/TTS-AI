package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func startDB() (*dynamodb.DynamoDB) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}));

    // Create DynamoDB client
	client := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")});
    
    if client == nil {
        return nil
    }

	return client
}