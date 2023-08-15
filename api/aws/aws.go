package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type OvertoneProps struct {
	awscdk.StackProps
}

func OvertoneStack(scope constructs.Construct, id string, props *OvertoneProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)


	// create AmazonDynamoDBFullAccess role
	dynamoDBRole := awsiam.NewRole(stack, aws.String("myDynamoDBFullAccessRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(aws.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromManagedPolicyArn(stack, aws.String("AmazonDynamoDBFullAccess"), aws.String("arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess")),
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")),
		},
	})


	// Create DynamoDB table
	awsdynamodb.NewTable(stack, jsii.String("OvertoneTable"), &awsdynamodb.TableProps{
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   jsii.String("OvertoneTable"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: aws.String("UserId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	// Create the Lambda function for createUser
	createUserLambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("createUserLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/createUser"),  // Pointing to the lambda folder
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole,
	})

	// Create the Lambda function for getUser
	getUserLambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("getUserLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/getUser"),  // Pointing to the lambda folder with the getUser code
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole,
	})

	// Create the Lambda function for PollySynthesize
	pollySynthesizeLambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("pollySynthesizeLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/polly"), // Pointing to the lambda folder with the pollySynthesize code
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole, // Assuming you want to use the same role
	})

	// Create API Gateway
	restApi := awsapigateway.NewRestApi(stack, jsii.String("OvertoneAPI"), &awsapigateway.RestApiProps{
		
		RestApiName:    jsii.String("OvertoneAPI"),
		CloudWatchRole: jsii.Bool(false),
	})

	// Add a POST endpoint to create users
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
	pollyResource := restApi.Root().AddResource(jsii.String("synthesize"), nil)
	pollyResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(pollySynthesizeLambda, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
	)

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
