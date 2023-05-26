package processor

import "Uala/go-workshop/internal/repository"

// processor.go is going to contain logic before connecting to the database

type Processor interface {
	Setup()
	Process() error
}

type LambdaProcessor struct {
	//all dependencies here (logger, db, etc)
	ContactRepository repository.Repository
}

func (p *LambdaProcessor) Setup() {
	//Setup the dynamodb service
	//Singleton pattern
	if p.ContactRepository == nil {
		p.ContactRepository = &repository.LambdaRepository{}
	}
	p.ContactRepository.Setup()
}

func (p *LambdaProcessor) Process() error {
	//Functionality here, call to db, etc
	//called from main.go

	return nil
}
