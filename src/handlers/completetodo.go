package main

import (
	"fmt"
	"context"
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

var ddb *dynamodb.DynamoDB
func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}


func CompleteTodo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse id from request body
	id := request.PathParameters["id"]

	// TOFIX: not working
	// Write to DynamoDB
	tableName := aws.String(os.Getenv("TODOS_TABLE_NAME"))
	_, err := ddb.UpdateItem(&dynamodb.UpdateItemInput{
		// TODO: Expression attribute names
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":d": {
				BOOL: aws.Bool(true),
			},
		},
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		UpdateExpression: aws.String("set todo.done = :d"),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body: err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body: request.Body,
			StatusCode: 200,
		}, nil
	}
}

func main() {
	lambda.Start(CompleteTodo)
}