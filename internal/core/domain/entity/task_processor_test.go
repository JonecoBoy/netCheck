package entity

//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
//	"testing"
//	"time"
//)
//
//func TestProcessFixedTasks(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//	defer mt.Close()
//
//	mt.Run("process fixed tasks", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateCursorResponse(1, "cron_db.tasks", mtest.FirstBatch, bson.D{
//			{"_id", "task1"},
//			{"type", "fixed"},
//			{"time", time.Now().Format("15:04")},
//		}))
//		processFixedTasks()
//		// Add assertions as needed
//	})
//}
//
//func TestProcessIntervalTasks(t *testing.T) {
//	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
//	defer mt.Close()
//
//	mt.Run("process interval tasks", func(mt *mtest.T) {
//		mt.AddMockResponses(mtest.CreateCursorResponse(1, "cron_db.tasks", mtest.FirstBatch, bson.D{
//			{"_id", "task1"},
//			{"type", "interval"},
//			{"last_run", time.Now().Add(-2 * time.Hour)},
//		}))
//		processIntervalTasks()
//		// Add assertions as needed
//	})
//}
