package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connects MongoDB
func ConnectDB() *mongo.Client {
	url := GetMongoURL()
	if url == "" {
		log.Fatalln("Error in loading env", url)
		return nil
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Returns client
var DB *mongo.Client = ConnectDB()

// Returns collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("AuthService").Collection(collectionName)
	return collection
}
