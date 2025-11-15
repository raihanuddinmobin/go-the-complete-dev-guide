package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"mobin.dev/pkg/config"
)

var client *mongo.Client

func Connect() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cnf := config.Get()

	if cnf.MongoURI == "" {
		return nil, fmt.Errorf("Mongo URI is Empty")
	}

	clientOptions := options.Client().ApplyURI(cnf.MongoURI).SetMaxPoolSize(50).SetMinPoolSize(5).SetMaxConnIdleTime(5 * time.Minute)

	c, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to connect mongo : %w ", err)
	}

	if err = c.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo : %w ", err)
	}

	client = c
	fmt.Println("Successfully Connected MongoDB ✅")
	return c, nil
}

func Disconnect() {
	if client != nil {
		_ = client.Disconnect(context.Background())
		log.Println("MongoDB disconnected ⛔")
	}
}
