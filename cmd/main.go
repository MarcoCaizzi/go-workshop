package main

import (
	"Uala/go-workshop/pkg/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := handler.LambdaHandler{} //This is the dependency injection

	lambda.Start(h.Create) //This is the entry point of the lambda function
}
