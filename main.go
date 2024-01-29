package main

import (
	"context"

)


func main() {
    ctx := context.Background()
    dynamodbRepo := NewDynamoDBRepo()
    itemService := NewItemService(dynamodbRepo)
    itemService.Add(ctx, Item{
        ID: 5,
        Name: "Peperroni",
        Description: "Topping",
        Price: "5.99",
    })
}
