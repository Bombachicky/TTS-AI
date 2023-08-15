package lib

import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
    "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
    "github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go/aws"
)

func CreateAPI(stack awscdk.Stack, createUserLambda awslambda.Function, getUserLambda awslambda.Function , pollySynthesizeLambda awslambda.Function) awsapigateway.RestApi {
    
	// Create API Gateway
	restApi := awsapigateway.NewRestApi(stack, jsii.String("OvertoneAPI"), &awsapigateway.RestApiProps{
		
		RestApiName:    jsii.String("OvertoneAPI"),
		CloudWatchRole: jsii.Bool(false),
	})

	userResource := restApi.Root().AddResource(jsii.String("users"), nil)
	userResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(createUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
	)

	// Add a GET endpoint to get a user by UserId
	// Add a new resource for the {userId} path parameter under /users
	userIdResource := userResource.AddResource(jsii.String("{userId}"), nil)
	userIdResource.AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getUserLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
	)

	// Add a POST endpoint to synthesize text to speech
	pollyResource := userResource.AddResource(jsii.String("synthesize"), nil)
	pollyResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(pollySynthesizeLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
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
