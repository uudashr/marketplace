package mongo

import (
	"context"

	"github.com/uudashr/marketplace/internal/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CategoryRepository is repository for product category.
type CategoryRepository struct {
	db *mongo.Database
}

// NewCategoryRepository constructs new product category repository.
func NewCategoryRepository(db *mongo.Database) (*CategoryRepository, error) {
	return &CategoryRepository{
		db: db,
	}, nil
}

// Store stores the product category.
func (r *CategoryRepository) Store(cat *product.Category) error {
	_, err := r.db.Collection("categories").InsertOne(context.TODO(), buildCategoryDoc(cat))

	// TODO: how to handle unique name
	if err != nil {
		return err
	}

	return nil
}

// CategoryByID retrieves product category by ID.
func (r *CategoryRepository) CategoryByID(id string) (*product.Category, error) {
	res := r.db.Collection("categories").FindOne(context.TODO(), bson.M{"_id": id})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var doc categoryDoc
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.build()
}

// Categories retrieves product categories.
func (r *CategoryRepository) Categories() ([]*product.Category, error) {
	cur, err := r.db.Collection("categories").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(context.TODO())
	}()

	var out []*product.Category
	for cur.Next(context.TODO()) {
		var doc categoryDoc
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

type categoryDoc struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (doc categoryDoc) build() (*product.Category, error) {
	return product.NewCategory(doc.ID, doc.Name)
}

func buildCategoryDoc(cat *product.Category) categoryDoc {
	return categoryDoc{
		ID:   cat.ID(),
		Name: cat.Name(),
	}
}
