package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection will have the MongoDB collection
var Collection *mongo.Collection

// GetCount returns the count of each topic
func GetCount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	topic, ok := req.URL.Query()["topic"]

	if !ok {
		log.Fatal("Error in request for getCount()")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{"topic": topic[0]}
	result, err := Collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal("Error while counting documents", err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
	cancel()
}
