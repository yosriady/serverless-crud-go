package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/satori/go.uuid"
)

type Todo struct {
	ID          string  `json:"id"`
	Description string 	`json:"description"`
	Done        bool   	`json:"done"`
	CreatedAt   string 	`json:"created_at"`
}

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

func AddTodo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("AddTodo")

	var (
		id = uuid.Must(uuid.NewV4(), nil).String()
		tableName = aws.String(os.Getenv("TODOS_TABLE_NAME"))
	)

	// Initialize todo
	todo := &Todo{
		ID:					id,
		Done:				false,
		CreatedAt:			time.Now().String(),
	}

	// Parse request body
	json.Unmarshal([]byte(request.Body), todo)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(todo)
	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: tableName,
	}
	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body: err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		body, _ := json.Marshal(todo)
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body: string(body),
			StatusCode: 200,
		}, nil
	}
}

func main() {
	lambda.Start(AddTodo)
}