package mongodb

import (
	"context"
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type MongoTaskRepository struct {
	Collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *MongoTaskRepository {
	return &MongoTaskRepository{Collection: collection}
}

func (r *MongoTaskRepository) Save(task *entity.Task) error {
	_, err := r.Collection.InsertOne(context.TODO(), task)
	return err
}

func (r *MongoTaskRepository) FindFixedTasks(now time.Time) ([]*entity.Task, error) {
	filter := bson.M{
		"type": "fixed",
		"time": now.Format("15:04"),
	}
	cursor, err := r.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []*entity.Task
	for cursor.Next(context.TODO()) {
		var task entity.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Decode error:", err)
			continue
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *MongoTaskRepository) FindIntervalTasks(now time.Time) ([]*entity.Task, error) {
	filter := bson.M{
		"type": "interval",
		"$or": []bson.M{
			{"last_run": bson.M{"$exists": false}},
			{"last_run": bson.M{"$lt": now.Add(-time.Second)}},
		},
	}
	cursor, err := r.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []*entity.Task
	for cursor.Next(context.TODO()) {
		var task entity.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Decode error:", err)
			continue
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *MongoTaskRepository) DeleteByID(id pkgEntity.ID) error {
	_, err := r.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *MongoTaskRepository) Update(task *entity.Task) error {
	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": task.ID},
		bson.M{
			"$set": bson.M{
				"name":               task.Name,
				"description":        task.Description,
				"type":               task.Type,
				"scheduled_time":     task.ScheduledTime,
				"scheduled_interval": task.ScheduledInterval,
				"cron_tab":           task.CronTab,
				"command":            task.Command,
				"updated_at":         time.Now(),
			},
		},
	)
	return err
}

func (r *MongoTaskRepository) UpdateLastRun(id pkgEntity.ID) error {
	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"last_run": time.Now()}},
	)
	return err
}
