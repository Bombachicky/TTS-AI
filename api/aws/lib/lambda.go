package lib
import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
    "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
    "github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func CreateUserLambda(stack awscdk.Stack, dynamoDBRole awsiam.Role) awslambda.Function {
    return awscdklambdagoalpha.NewGoFunction(stack, jsii.String("createUserLambda"), &awscdklambdagoalpha.GoFunctionProps{
        Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/createUser"),  // Pointing to the lambda folder
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole,
    })
}

func GetUserLambda(stack awscdk.Stack, dynamoDBRole awsiam.Role) awslambda.Function {
    return awscdklambdagoalpha.NewGoFunction(stack, jsii.String("getUserLambda"), &awscdklambdagoalpha.GoFunctionProps{
        Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/getUser"),  // Pointing to the lambda folder with the getUser code
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole,
    })
}

func CreatePollySynthesizeLambda(stack awscdk.Stack, dynamoDBRole awsiam.IRole) awslambda.Function {
	return awscdklambdagoalpha.NewGoFunction(stack, jsii.String("pollySynthesizeLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/polly"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole,
	})
}

func CreateOpenAILambda(stack awscdk.Stack, dynamoDBRole awsiam.IRole) awslambda.Function {
	return awscdklambdagoalpha.NewGoFunction(stack, jsii.String("openAILambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambda/openai"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Role: dynamoDBRole,
	})
}