package lib

import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
    "github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func CreateUserPool(stack awscdk.Stack) awscognito.UserPool {
	return awscognito.NewUserPool(stack, jsii.String("OvertoneUserPool"), &awscognito.UserPoolProps{
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
			// Add other desired properties...
		})
}