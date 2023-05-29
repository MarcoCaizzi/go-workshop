package main

import (
	"Uala/go-workshop/internal/processor"
	"Uala/go-workshop/pkg/handler"
	ddb "github.com/Bancar/uala-platform-go-utils/uala-platform-dynamodb-repository/pkg/search"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	r := ddb.NewContactsRepository()
	p := processor.NewProcessor(r)
	h := handler.NewHandler(p) //This is the dependency injection

	lambda.Start(h.Create) //This is the entry point of the lambda function
}
