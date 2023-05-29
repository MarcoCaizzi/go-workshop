package handler

import (
	"Uala/go-workshop/internal/processor"
	"Uala/go-workshop/pkg/dto"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

// Is in charge of receiving the request and forwarding it to the processor.go and its response
// is the response type for the lambda function
// Validating the input
// Validating the output

type (
	Response = events.APIGatewayProxyResponse
)

type Handler interface {
	Create(ctx context.Context, req dto.Request) (Response, error)
}

type LambdaHandler struct {
	//all dependencies here (logger, db, etc)
	ContactProcessor processor.Processor
}

func NewHandler(p processor.Processor) Handler {
	return &LambdaHandler{
		ContactProcessor: p,
	}
}

func (h *LambdaHandler) Create(ctx context.Context, req dto.Request) (Response, error) {
	//Functionality here, call to db, etc.
	//called from main.go
	if err := validateRequest(req); err != nil {
		lambdaError := dto.LambdaError{
			Code: dto.ValidationErrorCode,
			Msg:  err.Error(),
		}

		return Response{
			StatusCode: http.StatusBadRequest,
			Body:       lambdaError.Error(),
		}, nil
	}
	contact, err := h.ContactProcessor.Process(req)
	if err != nil {
		lambdaError := dto.LambdaError{
			Code: dto.InternalServerErrorCode,
			Msg:  err.Error(),
		}

		return Response{
			StatusCode: http.StatusInternalServerError,
			Body:       lambdaError.Error(),
		}, nil
	}
	return Response{
		StatusCode: http.StatusOK,
		Body:       contact.ToJsonStr(),
	}, nil
}

func validateRequest(req dto.Request) error {
	if req.FirstName == "" {
		return &dto.ValidationError{
			Field: "first_name",
			Err:   dto.WrongRequestError,
		}
	}
	if req.LastName == "" {
		return &dto.ValidationError{
			Field: "last_name",
			Err:   dto.WrongRequestError,
		}
	}
	return nil
}
