package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
)

func CompleteTodo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// 1. Parse id from request body
	// 2. ddb.UpdateItem to mark the row as Done: true
	// 3. Return HTTP response

	return events.APIGatewayProxyResponse{
		Body: request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(CompleteTodo)
}