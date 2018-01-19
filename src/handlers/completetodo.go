package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

func CompleteTodo(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("[completetodo] Received body: ", request.Body)
	return events.APIGatewayProxyResponse{
		Body: request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(CompleteTodo)
}