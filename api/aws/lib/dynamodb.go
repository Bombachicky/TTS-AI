package lib

import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
    "github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go/aws"
	
)

func CreateDynamoTable(stack awscdk.Stack) awsdynamodb.Table {
    return awsdynamodb.NewTable(stack, jsii.String("OverToneTable"), &awsdynamodb.TableProps{
        BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   jsii.String("OverToneTable"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: aws.String("Email"),
			Type: awsdynamodb.AttributeType_STRING,
		},
    })
}
