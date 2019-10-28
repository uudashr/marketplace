package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/uudashr/marketplace/internal/mongo/migrations"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	migrate "github.com/xakep666/mongo-migrate"
)

func main() {
	flagAction := flag.String("action", "", "The action ('up' or 'down')")
	flagDBAddress := flag.String("db-address", "localhost:27017", "Database address")
	flagDBName := flag.String("db-name", "marketplace_test", "Database name")
	flagHelp := flag.Bool("help", false, "Show help")

	flag.Parse()
	if *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *flagAction == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := openDB(ctx, *flagDBAddress, *flagDBName)
	if err != nil {
		panic(fmt.Errorf("Fail to open db: %w", err))
	}
	defer func() {
		if err = db.Client().Disconnect(context.TODO()); err != nil {
			log.Println("Fail to disconnect from mongodb:", err)
		}
	}()

	migrate.SetDatabase(db)

	switch *flagAction {
	case "up":
		if err := migrate.Up(migrate.AllAvailable); err != nil {
			panic(err)
		}
	case "down":
		if err := migrate.Down(migrate.AllAvailable); err != nil {
			panic(err)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func openDB(ctx context.Context, dbAddress, dbName string) (*mongo.Database, error) {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", dbAddress))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
