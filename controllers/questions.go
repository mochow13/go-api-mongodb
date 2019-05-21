package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/the-flying-dutchman/prostuti-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection will have the MongoDB collection
var Collection *mongo.Collection

const moduleSize int = 50

// GetCount returns the count of each topic
func GetCount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	topic, ok := req.URL.Query()["topic"]

	if !ok {
		log.Fatal("Error in request for GetCount()")
	}

	filter := bson.M{"topic": topic[0]}
	result, err := Collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error while counting documents", err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

// GetQuestion returns questions for a topic
func GetQuestion(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	// get the query params
	topic, _ := req.URL.Query()["topic"]
	module, ok := req.URL.Query()["module"]
	moduleNum, _ := strconv.Atoi(module[0])

	if !ok {
		log.Fatal("Error in request for GetQuestions()")
	}

	// set query options
	var opts = new(options.FindOptions)
	opts = opts.SetSkip(int64(moduleNum * moduleSize))
	opts = opts.SetLimit(int64(moduleSize))
	filter := bson.M{"topic": topic[0]}

	cursor, err := Collection.Find(ctx, filter, opts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	defer cursor.Close(ctx)

	var questions []models.Question
	for cursor.Next(ctx) {
		var ques models.Question
		cursor.Decode(&ques)
		questions = append(questions, ques)
	}

	if err := cursor.Err(); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(questions)
}
