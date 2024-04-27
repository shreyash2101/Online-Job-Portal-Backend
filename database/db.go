package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://shreyash2101:<password>@cluster0.wvu0mzd.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "jobPortal"
const colName = "jobs"

var Collection *mongo.Collection

// connect to mongodb

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb client
	client,err := mongo.Connect(context.TODO(), clientOption)

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	// collection instance
	Collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}