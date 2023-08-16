package lib

import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
    "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
    "github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go/aws"
)

// Create API Func
// Parameters: createUserLambda, getUserLambda, pollySynthesizeLambda
// Return Type: awsapigateway.RestApi 
func CreateAPI(stack awscdk.Stack, createUserLambda awslambda.Function, getUserLambda awslambda.Function , pollySynthesizeLambda awslambda.Function ) awsapigateway.RestApi {
    
	// Create API Gateway
	restApi := awsapigateway.NewRestApi(stack, jsii.String("OvertoneAPI"), &awsapigateway.RestApiProps{
		
		RestApiName:    jsii.String("OvertoneAPI"),
		CloudWatchRole: jsii.Bool(false),
	})

	// Create Cognito User Pool With Required Attributes : Email And Auto Email Verification
	userPool := awscognito.NewUserPool(stack, jsii.String("OvertoneUserPool"), &awscognito.UserPoolProps{
		UserPoolName: jsii.String("OvertoneUserPool"),
		SignInAliases: &awscognito.SignInAliases{
			Username: jsii.Bool(true),
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

	// Sign In Page EndPoint With Cognito Authorizer: GET
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

	// Add Cors Prefight configuration
	userResource.AddCorsPreflight(&awsapigateway.CorsOptions{
		AllowHeaders: jsii.Strings("Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent"),
		AllowMethods: jsii.Strings("OPTIONS,POST,GET"),
		AllowOrigins: jsii.Strings("*"), // For development, you can limit this to your localhost
		MaxAge:       awscdk.Duration_Seconds(aws.Float64(3600)),
	})

		// log lambda function ARN
	awscdk.NewCfnOutput(stack, jsii.String("lambdaFunctionArn"), &awscdk.CfnOutputProps{
		Value:       createUserLambda.FunctionArn(),
		Description: jsii.String("Lambda function ARN"),
	})

	
	
    return restApi
}
