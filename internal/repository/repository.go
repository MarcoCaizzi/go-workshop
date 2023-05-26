package repository

// The repository.go its function will be to interact with dynamoDB
import (
	"Uala/go-workshop/pkg/dto"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	TableName = "Contacts"
)

type Repository interface {
	// Setup all methods here
	Setup()
	Insert(contact dto.Contact) (dto.Contact, error)
}

type LambdaRepository struct {
	TableName string             //This is the dependency injection for the table name
	svc       *dynamodb.DynamoDB //This is the dependency injection for the dynamodb service
}

func (r *LambdaRepository) Setup() {
	//Setup the dynamodb service
	r.TableName = TableName //Instance the table name

	//Create a new session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	r.svc = dynamodb.New(sess)
}

func (r *LambdaRepository) Insert(contact dto.Contact) (dto.Contact, error) {
	//Insert the contact in dynamodb
	//Convert the Record Go type to dynamodb attribute value type using MarshalMap
	//https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.MarshalMap
	//https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/#MarshalMap
	item, err := dynamodbattribute.MarshalMap(contact)

	// Manage the error
	if err != nil {
		return dto.Contact{}, err
	}

	// Declare a new PutItemInput instance
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: &r.TableName, //aws.String(r.TableName) --> This is the same
	}

	// Put the item in the table
	_, err = r.svc.PutItem(input)

	// Manage the error
	if err != nil {
		return dto.Contact{}, err
	}
	return contact, nil
}
