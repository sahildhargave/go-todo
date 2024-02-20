//☮✝☪🕉☸🕉☪✝

package middleware

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadEnv()
	createDBInstance()

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectingString := os.Getenv("MONGO_URL")
	mongo_db_name := os.Getenv("MONGO_DB_NAME")
	collName := os.Getenv("MONGO_DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectingString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")
	collection = client.Database(mongo_db_name).Collection(collName)
	fmt.Println("Collection instance created")
}
