package tasks

import (
	"context"
	"fmt"
	"github.com/JonecoBoy/netCheck/internal/utils/config"
	"github.com/JonecoBoy/netCheck/pkg/entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartTaskProcessor() {
	go processFixedTasks()
	go processIntervalTasks()
}

// Process fixed-time tasks
func processFixedTasks() {
	collection := config.MongoClient.Database("cron_db").Collection("tasks")

	for {
		now := time.Now().Format("15:04")
		filter := bson.M{"type": "fixed", "time": now}

		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Println("Error fetching fixed tasks:", err)
			continue
		}

		for cursor.Next(context.TODO()) {
			var task Task
			if err := cursor.Decode(&task); err != nil {
				log.Println("Error decoding task:", err)
				continue
			}
			executeTask(task)
		}

		time.Sleep(60 * time.Second)
	}
}

// Process interval-based tasks
func processIntervalTasks() {
	collection := config.MongoClient.Database("cron_db").Collection("tasks")

	for {
		now := time.Now()
		filter := bson.M{
			"type": "interval",
			"$or": []bson.M{
				{"last_run": bson.M{"$exists": false}},
				{"last_run": bson.M{"$lt": now.Add(-time.Second)}},
			},
		}

		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Println("Error fetching interval tasks:", err)
			continue
		}

		for cursor.Next(context.TODO()) {
			var task Task
			if err := cursor.Decode(&task); err != nil {
				log.Println("Error decoding task:", err)
				continue
			}

			if time.Since(*task.LastRun).Seconds() >= float64(*task.ScheduledInterval) {
				executeTask(task)
				updateTaskLastRun(collection, task.ID)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

// Execute task
func executeTask(task Task) {
	fmt.Println("Executing task:", task.Command)
}

// Update last run time
func updateTaskLastRun(collection *mongo.Collection, taskID entity.ID) {
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": taskID},
		bson.M{"$set": bson.M{"last_run": time.Now()}},
	)
	if err != nil {
		log.Println("Error updating last run:", err)
	}
}
