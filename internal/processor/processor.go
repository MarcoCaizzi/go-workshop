package processor

import (
	"Uala/go-workshop/internal/repository"
	"Uala/go-workshop/pkg/dto"
)

// processor.go is going to contain logic before connecting to the database

type Processor interface {
	Process(req dto.Request) (dto.Contact, error) //This is the entry point of the lambda function
}

type LambdaProcessor struct {
	//all dependencies here (logger, db, etc)
	ContactRepository repository.Repository
}

func NewProcessor(r repository.Repository) *LambdaProcessor {
	return &LambdaProcessor{ContactRepository: r}
}

func (p *LambdaProcessor) Process(req dto.Request) (dto.Contact, error) {
	//Functionality here, call to db, etc.
	//called from main.go
	contact := dto.Contact{
		ID:        "",
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}
	item, err := p.ContactRepository.Insert(contact)
	if err != nil {
		return dto.Contact{}, err
	}
	return item, nil
}
