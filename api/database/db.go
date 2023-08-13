package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}));
    
    // Create DynamoDB client
	client := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")});
    
    if client == nil {
        fmt.Println("Could not create DynamoDB client")
        return
    }

	// Create the input configuration instance
    input := &dynamodb.ListTablesInput{};

    fmt.Printf("Tables:\n")

    for {
        // Get the list of tables
        result, err := client.ListTables(input)
        if err != nil {
            if aerr, ok := err.(awserr.Error); ok {
                switch aerr.Code() {
                case dynamodb.ErrCodeInternalServerError:
                    fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
                default:
                    fmt.Println(aerr.Error())
                }
            } else {
                // Print the error, cast err to awserr.Error to get the Code and
                // Message from an error.
                fmt.Println(err.Error())
            }
            return
        }

        for _, n := range result.TableNames {
            fmt.Println(*n)
        }

        // assign the last read tablename as the start for our next call to the ListTables function
        // the maximum number of table names returned in a call is 100 (default), which requires us to make
        // multiple calls to the ListTables function to retrieve all table names
        input.ExclusiveStartTableName = result.LastEvaluatedTableName

        if result.LastEvaluatedTableName == nil {
            break
        }
    }
}