package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.MustRegister(func(db *mongo.Database) error {
		opts := options.Index().SetName("ix_strcat")
		keys := bson.M{"storeId": 1, "categoryId": 1}
		model := mongo.IndexModel{Keys: keys, Options: opts}

		_, err := db.Collection("products").Indexes().CreateOne(context.TODO(), model)
		return err
	}, func(db *mongo.Database) error {
		return db.Collection("products").Drop(context.TODO())
	})
}
