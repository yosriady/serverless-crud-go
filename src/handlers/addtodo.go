package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
	"os"
)

func AddTodo(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// TODO: access environment variable
	fmt.Printf(os.Getenv("TODOS_TABLE_NAME"))

	// TODO: use aws sdk to connect to dynamoDB

	// TODO: parse request body and write to dynamoDB

	fmt.Println("[addtodo] Received body: ", request.Body)
	return events.APIGatewayProxyResponse{
		Body: request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(AddTodo)
}