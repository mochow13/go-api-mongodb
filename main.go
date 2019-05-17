package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/the-flying-dutchman/prostuti-api/controllers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration is definition of envs
type Configuration struct {
	DBName     string
	DBUserName string
	DBPassword string
	ModuleSize int
}

func buildURI() string {
	var file, _ = os.Open("env.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error occurred while reading envs", err)
	}
	// fmt.Println(configuration.DBName, configuration.DBPassword)
	mongoURI := "mongodb://" + config.DBUserName + ":" + config.DBPassword + "@ds155252.mlab.com:55252/" + config.DBName
	fmt.Println(mongoURI)
	return mongoURI
}

func configDB(mongoURI string) *mongo.Client {
	fmt.Println("Starting connection...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error while connecting ", err)
	}
	err = client.Ping(ctx, nil)
	fmt.Println("Connected to MongoDB!")
	return client
}

func defineRoutes(client *mongo.Client) (router *mux.Router) {
	router = mux.NewRouter()
	collection := client.Database("memento").Collection("bcs_questions")
	controllers.Collection = collection
	router.HandleFunc("/v1/questions/getcount", controllers.GetCount).Methods("GET")
	return
}

func main() {
	mongoURI := buildURI()
	client := configDB(mongoURI)
	router := defineRoutes(client)
	log.Fatal(http.ListenAndServe(":8000", router))
}
