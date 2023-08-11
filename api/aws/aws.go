package main

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AwsStackProps struct {
	awscdk.StackProps
}
type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewAwsStack(scope constructs.Construct, id string, props *AwsStackProps, table *TableBasics) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	

	// awsdynamodb.NewTable(stack, jsii.String("Test-Table"), &awsdynamodb.TableProps{
	// 	PartitionKey: &awsdynamodb.Attribute{
	// 		Name: jsii.String("userId"),
	// 		Type: awsdynamodb.AttributeType_STRING,
	// 	},
	// 	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	// })

	return stack

}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	stackProps := &awscdk.StackProps{
		Env: env(),
	}

	stack := awscdk.NewStack(app, jsii.String("aws"), stackProps)

	awsdynamodb.NewTable(stack, jsii.String("Test-Table"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	//return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("533105544577"),
		Region:  jsii.String("us-east-1"),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
