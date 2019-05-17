package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	fmt.Println("Starting connection...")
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		fmt.Println("Error occurred", err)
	}
	client, err_connect := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err_connect != nil {
		fmt.Println("Error while connecting to mongodb...", err)
	}
}
