package provider

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProvider struct {
	client *mongo.Client
}

func (p *MongoProvider) Connect(uri string) error {
	var err error
	p.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	return err
}

func (p *MongoProvider) Disconnect() error {
	return p.client.Disconnect(context.Background())
}
