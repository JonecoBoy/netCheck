package entity

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func runFixedTaskCycle(collection *mongo.Collection, now string) {
	filter := bson.M{"type": "fixed", "time": now}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error fetching fixed tasks:", err)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Error decoding task:", err)
			continue
		}
		executeTask(task)
	}
}
