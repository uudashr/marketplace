// +build integration

package mongo_test

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	_ "github.com/uudashr/marketplace/internal/mongo/migrations"
	migrate "github.com/xakep666/mongo-migrate"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbAddress = flag.String("db-address", "localhost:27017", "Database address")
	dbName    = flag.String("db-name", "marketplace_test", "Database name")
)

type dbFixture struct {
	t  *testing.T
	db *mongo.Database
}

func (s *dbFixture) tearDown() {
	if err := migrate.Down(migrate.AllAvailable); err != nil {
		s.t.Error("fail to migrate down:", err)
	}

	if err := s.db.Client().Disconnect(context.TODO()); err != nil {
		s.t.Error("fail to disconnect client:", err)
	}
}

func setupDBFixture(t *testing.T) *dbFixture {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", *dbAddress))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		t.Fatal("err:", err)
	}

	db := client.Database(*dbName)
	migrate.SetDatabase(db)
	if err = migrate.Up(migrate.AllAvailable); err != nil {
		t.Fatal("err:", err)
	}

	return &dbFixture{
		t:  t,
		db: db,
	}
}
