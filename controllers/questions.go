package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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

	filter := bson.M{"topic": topic}
	result, err := Collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error while counting documents", err)
	}
	json.NewEncoder(res).Encode(result)
}
