package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
)

func ListTodos(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// 1. Call ddb scan
	// 2. Return

	return events.APIGatewayProxyResponse{
		Body: request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(ListTodos)
}