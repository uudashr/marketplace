package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/shopspring/decimal"

	"github.com/uudashr/marketplace/internal/product"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository is repository for Product.
type ProductRepository struct {
	db *mongo.Database
}

// NewProductRepository constructs new product repository.
func NewProductRepository(db *mongo.Database) (*ProductRepository, error) {
	return &ProductRepository{
		db: db,
	}, nil
}

// Store stores product.
func (r *ProductRepository) Store(prd *product.Product) error {
	doc, err := buildProductDoc(prd)
	if err != nil {
		return err
	}

	_, err = r.db.Collection("products").InsertOne(context.TODO(), doc)

	// TODO: how to handle unique name
	if err != nil {
		return err
	}

	return nil
}

// ProductByID on the repository.
func (r *ProductRepository) ProductByID(id string) (*product.Product, error) {
	res := r.db.Collection("products").FindOne(context.TODO(), bson.M{"_id": id})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var doc productDoc
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.build()
}

// Products retrieves products.
func (r *ProductRepository) Products() ([]*product.Product, error) {
	cur, err := r.db.Collection("products").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cur.Close(context.TODO())
	}()

	var out []*product.Product
	for cur.Next(context.TODO()) {
		var doc productDoc
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

type productDoc struct {
	ID          string               `bson:"_id"`
	StoreID     string               `bson:"storeId"`
	CategoryID  string               `bson:"categoryId"`
	Name        string               `bson:"name"`
	Price       primitive.Decimal128 `bson:"price"`
	Description string               `bson:"description"`
	Quantity    int                  `bson:"quantity"`
}

func (doc productDoc) build() (*product.Product, error) {
	price, err := decimal.NewFromString(doc.Price.String())
	if err != nil {
		return nil, err
	}

	return product.New(doc.ID, doc.StoreID, doc.CategoryID, doc.Name, price, doc.Description, doc.Quantity)
}

func buildProductDoc(prd *product.Product) (productDoc, error) {
	price, err := primitive.ParseDecimal128(prd.Price().String())
	if err != nil {
		return productDoc{}, err
	}

	return productDoc{
		ID:          prd.ID(),
		StoreID:     prd.StoreID(),
		CategoryID:  prd.CategoryID(),
		Name:        prd.Name(),
		Price:       price,
		Description: prd.Description(),
		Quantity:    prd.Quantity(),
	}, nil
}
