package main

import (
	"fmt"
	"context"
  "github.com/aws/aws-lambda-go/lambda"
)

func hello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("hello %s", name), nil
}

func main() {
	lambda.Start(hello)
}
