package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

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

var client *mongo.Client
var mongoURI string

func buildURI() {
	var file, _ = os.Open("env.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error occurred while reading envs", err)
	}
	// fmt.Println(configuration.DBName, configuration.DBPassword)
	mongoURI = "mongodb://" + config.DBUserName + ":" + config.DBPassword + "@ds155252.mlab.com:55252/" + config.DBName
	fmt.Println(mongoURI)
}

func configDB() {
	fmt.Println("Starting connection...")
	// ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	// if err != nil {
	// 	log.Fatal("Error occurred ", err)
	// }
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, errConnect := mongo.Connect(context.TODO(), clientOptions)
	if errConnect != nil {
		log.Fatal(errConnect)
	}
	errConnect = client.Ping(context.TODO(), nil)
	if errConnect != nil {
		log.Fatal(errConnect)
	}
	fmt.Println("Connected to MongoDB!")
}

func main() {
	buildURI()
	configDB()
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
