package mongo

import (
	"context"

	"github.com/uudashr/marketplace/internal/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// StoreRepository is repository for Store.
type StoreRepository struct {
	db *mongo.Database
}

// NewStoreRepository constructs new store repository.
func NewStoreRepository(db *mongo.Database) (*StoreRepository, error) {
	return &StoreRepository{
		db: db,
	}, nil
}

// Store stores/puts store.
func (r *StoreRepository) Store(str *store.Store) error {
	_, err := r.db.Collection("stores").InsertOne(context.TODO(), buildStoreDoc(str))

	// TODO: how to handle unique name
	if err != nil {
		return err
	}

	return nil
}

// StoreByID on the repository.
func (r *StoreRepository) StoreByID(id string) (*store.Store, error) {
	res := r.db.Collection("stores").FindOne(context.TODO(), bson.M{"_id": id})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var doc storeDoc
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.build()
}

// Stores retrieves stores.
func (r *StoreRepository) Stores() ([]*store.Store, error) {
	cur, err := r.db.Collection("stores").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(context.TODO())
	}()

	var out []*store.Store
	for cur.Next(context.TODO()) {
		var doc storeDoc
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}

		cat, err := doc.build()
		if err != nil {
			return nil, err
		}

		out = append(out, cat)

	}

	return out, nil
}

type storeDoc struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (doc storeDoc) build() (*store.Store, error) {
	return store.New(doc.ID, doc.Name)
}

func buildStoreDoc(str *store.Store) storeDoc {
	return storeDoc{
		ID:   str.ID(),
		Name: str.Name(),
	}
}
