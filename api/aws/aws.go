package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"aws/lib"  
)

type OvertoneProps struct {
	awscdk.StackProps
}

func OvertoneStack(scope constructs.Construct, id string, props *OvertoneProps) awscdk.Stack {
	// Stack Creation
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create DynamoDB Table
	 lib.CreateDynamoTable(stack)

	// Create IAM Role for Lambdas
	lambdaRole := lib.CreateLambdaIAMRole(stack)

	// Setup Lambda Functions
	createUserLambda := lib.CreateUserLambda(stack, lambdaRole)
	getUserLambda := lib.GetUserLambda(stack, lambdaRole)
	createPollySynthesizeLambda := lib.CreatePollySynthesizeLambda(stack, lambdaRole)
	createOpenAIResponseLambda := lib.CreateOpenAIResponseLambda(stack, lambdaRole)

	// Create API and Endpoints and Add Cognito Authorizer
	lib.CreateAPI(stack, createUserLambda, getUserLambda, createPollySynthesizeLambda, createOpenAIResponseLambda)
	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	OvertoneStack(app, "OvertoneStack", &OvertoneProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
