package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
    ID int 
    Name string
    Description string
    Price string
}


type ItemRepo interface {
    GetByID(ctx context.Context, ID int) (*Item, error)
    Add(ctx context.Context, i Item) error
}

type DynamoDBItemRepo struct {
   db *dynamodb.DynamoDB
}


func NewDynamoDBRepo() *DynamoDBItemRepo {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1")},
    )
    if err != nil {
        fmt.Printf("error %v", err)
    }
    return &DynamoDBItemRepo{
        db: dynamodb.New(sess),
    }
}


func (d *DynamoDBItemRepo) GetByID(ctx context.Context, ID int) (*Item, error) {
    return nil, nil
}

func (d *DynamoDBItemRepo) Add(ctx context.Context, i Item) error {
    av, err := dynamodbattribute.MarshalMap(i)
    if err != nil {
        log.Fatalf("Got error marshalling new item: %s", err)
    }
    tableName := "Inventory"

    input := &dynamodb.PutItemInput{
        Item: av,
        TableName: aws.String(tableName),
    }
    
    _, err = d.db.PutItem(input)
    if err != nil {
        log.Fatalf("Got error calling PutItem: %s", err)
    }

    fmt.Printf("Successfully added %s to table %s.", i.Name, tableName)
    return nil
}


// service definition
type ItemService struct {
    repo ItemRepo 
}

func NewItemService(repo ItemRepo) *ItemService {
    return &ItemService{
        repo: repo,
    }
}

func (is *ItemService) Add(ctx context.Context, i Item) error {
    // business logic

    err := is.repo.Add(ctx, i)
    if err != nil {
        return err
    }

    return nil
}


