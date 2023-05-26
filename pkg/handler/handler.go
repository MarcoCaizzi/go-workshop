package handler

import (
	"Uala/go-workshop/pkg/dto"
	"context"
	"github.com/aws/aws-lambda-go/events"
)

// Is in charge of receiving the request and forwarding it to the processor.go and its response
// Response is the response type for the lambda function
// Validating the input
// Validating the output

type Response = events.APIGatewayProxyResponse

type Handler interface {
	Create(ctx context.Context, req dto.Request) (Response, error)
}

type LambdaHandler struct {
	//all dependencies here (logger, db, etc)

}

func (h *LambdaHandler) Create(ctx context.Context, req dto.Request) (Response, error) {
	//Functionality here, call to db, etc
	//called from main.go

	return Response{}, nil
}
