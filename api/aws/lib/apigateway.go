package lib

import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
    "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
    "github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	
)

// Create API Func
// Parameters: createUserLambda, getUserLambda, pollySynthesizeLambda
// Return Type: awsapigateway.RestApi 
func CreateAPI(stack awscdk.Stack, createUserLambda awslambda.Function, getUserLambda awslambda.Function , pollySynthesizeLambda awslambda.Function, openAIResponseLambda awslambda.Function) awsapigateway.RestApi {
    
	// Create API Gateway
	restApi := awsapigateway.NewRestApi(stack, jsii.String("OvertoneAPI"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
		},
		RestApiName:    jsii.String("OvertoneAPI"),
		CloudWatchRole: jsii.Bool(true),
		DeployOptions: &awsapigateway.StageOptions{
		LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
		DataTraceEnabled: jsii.Bool(true),
	},
		
	})

	// Create Cognito User Pool With Required Attributes : Email And Auto Email Verification
	userPool := awscognito.NewUserPool(stack, jsii.String("UserPool"), &awscognito.UserPoolProps{
		UserPoolName: jsii.String("UserPool"),
		SignInAliases: &awscognito.SignInAliases{
			Email:    jsii.Bool(true),
		},
		StandardAttributes: &awscognito.StandardAttributes{
			Email: &awscognito.StandardAttribute{
				Mutable:   jsii.Bool(true),
				Required:  jsii.Bool(true),
			},
		},
		AutoVerify: &awscognito.AutoVerifiedAttrs{
		Email: jsii.Bool(true),
		},
		// Add other desired properties...
	})

	// User Pool Client
	awscognito.NewUserPoolClient(stack, jsii.String("MyAppClient"), &awscognito.UserPoolClientProps{
		UserPool: userPool,
		GenerateSecret: jsii.Bool(false), // Set to true if you need a secret (commonly for server-side apps)
		// ... other App Client configurations ...
	})

	// User Pool Authorization
	authorizer := awsapigateway.NewCognitoUserPoolsAuthorizer(stack, jsii.String("APIGatewayAuthorizer"), &awsapigateway.CognitoUserPoolsAuthorizerProps{
		CognitoUserPools: &[]awscognito.IUserPool{
		userPool,
		},
    })

	// Sign Up Page EndPoint With Cognito Authorizer: POST
	userResource := restApi.Root().AddResource(jsii.String("users"), nil)
	userResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(createUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			Authorizer:        authorizer,
			AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			
			
		},
	)

	// Enpoint to add users info to DB after cognito sign-up: POST
	userDbResource := userResource.AddResource(jsii.String("sign-up"), nil)
	userDbResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(createUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			// Authorizer:        authorizer,
			// AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			
		},
	)

	userSignIn:= userResource.AddResource(jsii.String("sign-in"), nil)
	userSignIn.AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			Authorizer:        authorizer,
			AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			
		},
	)

	// // Enpoint to add users info to DB: POST
	// userDbResource := userResource.AddResource(jsii.String("add"), nil)
	// userDbResource.AddMethod(
	// 	jsii.String("POST"),
	// 	awsapigateway.NewLambdaIntegration(createUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
	// 	&awsapigateway.MethodOptions{
	// 		// Authorizer:        authorizer,
	// 		// AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			
	// 	},
	// )


	//Sign In Page EndPoint With Cognito Authorizer: GET
	userIdResource := userResource.AddResource(jsii.String("{userId}"), nil)
	userIdResource.AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			Authorizer:        authorizer,
			AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			
		},
	)

	// Add a POST endpoint to synthesize text to speech
	pollyResource := userResource.AddResource(jsii.String("synthesize"), nil)
	pollyResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(pollySynthesizeLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			// Authorizer:        authorizer,
			// AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			
		},
	)

	// Add a POST endpoint to respond to user input
	openAIResource := userResource.AddResource(jsii.String("message"), nil)
	openAIResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(openAIResponseLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			// Authorizer:        authorizer,
			// AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
		},
	)

		// log lambda function ARN
	awscdk.NewCfnOutput(stack, jsii.String("lambdaFunctionArn"), &awscdk.CfnOutputProps{
		Value:       createUserLambda.FunctionArn(),
		Description: jsii.String("Lambda function ARN"),
	})

    return restApi
}
