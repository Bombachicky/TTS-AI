package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
)

type RequestInput struct {
	Text string `json:"text"`
}

type SynthesizedSpeechResponse struct {
	AudioStream []byte `json:"audioStream"`
}

var pollyClient *polly.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	pollyClient = polly.NewFromConfig(cfg)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var input RequestInput
	err := json.Unmarshal([]byte(request.Body), &input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error parsing request body",
			StatusCode: 400,
		}, err
	}

	resp, err := SynthesizeSpeech(ctx, input.Text)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error synthesizing speech",
			StatusCode: 500,
		}, err
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error marshalling response",
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func SynthesizeSpeech(ctx context.Context, text string) (SynthesizedSpeechResponse, error) {
	if text == "" {
		return SynthesizedSpeechResponse{}, errors.New("Text is required")
	}

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: types.OutputFormatMp3,
		Text:         &text,
		VoiceId:      types.VoiceIdJoanna,
	}

	output, err := pollyClient.SynthesizeSpeech(ctx, input)
	if err != nil {
		return SynthesizedSpeechResponse{}, err
	}

	buffer := make([]byte, 2048)
	stream := []byte{}
	for {
		numBytes, readErr := output.AudioStream.Read(buffer)
		if readErr != nil && readErr.Error() != "EOF" {
			return SynthesizedSpeechResponse{}, readErr
		}
		if numBytes == 0 {
			break
		}
		stream = append(stream, buffer[:numBytes]...)
	}

	return SynthesizedSpeechResponse{
		AudioStream: stream,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
