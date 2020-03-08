package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuestionsRepo - MonogoDB Connection
type QuestionsRepo struct {
	*mongo.Client
}

// SetupQuestionsRepo - Connect to Database and return connection
func SetupQuestionsRepo(ctx context.Context, config string) (*QuestionsRepo, error) {
	clientOptions := options.Client().ApplyURI(config)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return &QuestionsRepo{}, err
	}
	// Check the connection - Reduces Client resiliance...
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return &QuestionsRepo{}, err
	}

	return &QuestionsRepo{client}, nil
}
