package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// Access the configuration properties
	fmt.Println("Access Key ID:", cfg.Credentials.AccessKeyID)
	fmt.Println("Secret Access Key:", cfg.Credentials.SecretAccessKey)
	fmt.Println("Region:", cfg.Region)
}