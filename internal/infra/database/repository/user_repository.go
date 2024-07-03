package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/reangeline/micro_saas/internal/domain/entity"
)

type UserRepository struct {
	svc *dynamodb.DynamoDB
}

func NewUserRepository(svc *dynamodb.DynamoDB) *UserRepository {
	return &UserRepository{
		svc,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {

	usa, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	log.Printf("Marshalled User: %+v", usa)

	input := &dynamodb.PutItemInput{
		Item:      usa,
		TableName: aws.String("usersTable"),
	}

	_, err = ur.svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindAll() ([]*entity.User, error) {
	var userModel []*entity.User
	var userModelOut []*entity.User

	params := &dynamodb.ScanInput{
		TableName: aws.String("usersTable"),
	}

	result, err := ur.svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &userModel)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	userModelOut = append(userModelOut, userModel...)

	return userModelOut, nil
}

func (ur *UserRepository) FindByUserEmail(email string) (*entity.User, error) {
	var userModel *entity.User
	result, err := ur.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("usersTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &userModel)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return userModel, nil

}

func (ur *UserRepository) UpdateByEmail(input *entity.User) (*entity.User, error) {

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String("usersTable"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(input.Name),
			},
			":s": {
				S: aws.String(input.LastName),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#n": aws.String("name"),
		},
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(input.Email),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #n = :r, last_name = :s"),
	}

	_, err := ur.svc.UpdateItem(updateInput)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem xxxxx: %s", err)
	}

	return input, nil

}
