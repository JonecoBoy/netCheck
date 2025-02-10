package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() {
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(AppConfig.MongoURI))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}
	log.Println("Connected to MongoDB:", AppConfig.MongoURI)
}
