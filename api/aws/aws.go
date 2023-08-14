package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
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

	// Create the Lambda function for createUser
	createUserLambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("createUserLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/createUser"),  // Pointing to the lambda folder
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
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
