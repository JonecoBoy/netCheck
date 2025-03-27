package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{Collection: col}
}

func (r *MongoUserRepository) Create(user *entity.User) error {
	user.ID = pkgEntity.NewId()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := r.Collection.InsertOne(context.TODO(), user)
	return err
}

func (r *MongoUserRepository) Update(user *entity.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

func (r *MongoUserRepository) DeleteByID(id pkgEntity.ID) error {
	_, err := r.Collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func (r *MongoUserRepository) SoftDeleteByID(id pkgEntity.ID) error {
	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{"deleted_at": time.Now()}},
	)
	return err
}

func (r *MongoUserRepository) FindByID(id pkgEntity.ID) (*entity.User, error) {
	var user entity.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"id": id, "deleted_at": bson.M{"$eq": time.Time{}}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindAll(page, limit int) ([]*entity.User, error) {
	return r.Find(bson.M{}, page, limit)
}

func (r *MongoUserRepository) FindByRole(role int, page, limit int) ([]*entity.User, error) {
	return r.Find(bson.M{"role": role}, page, limit)
}

func (r *MongoUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"email": email, "deleted_at": bson.M{"$eq": time.Time{}}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) Find(filter bson.M, page, limit int) ([]*entity.User, error) {
	filter["deleted_at"] = bson.M{"$eq": time.Time{}}

	skip := int64((page - 1) * limit)
	cursor, err := r.Collection.Find(context.TODO(), filter, &options.FindOptions{
		Limit: &[]int64{int64(limit)}[0],
		Skip:  &skip,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*entity.User
	for cursor.Next(context.TODO()) {
		var user entity.User
		if err := cursor.Decode(&user); err == nil {
			users = append(users, &user)
		}
	}
	return users, nil
}
