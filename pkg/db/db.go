package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuestionsRepo - MonogoDB Connection
type QuestionsRepo struct {
	*mongo.Collection
}

// NewQuestionsRepo - Connect to Database and return connection
func NewQuestionsRepo(ctx context.Context, config string, db string) (*QuestionsRepo, error) {
	clientOptions := options.Client().ApplyURI(config)

	// TODO: Implement Initial Retry Logic Here Maybe? or higherlevel in main function?
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return &QuestionsRepo{}, err
	}
	// Check the connection - Reduces Client resiliance...
	err = client.Ping(ctx, nil)

	if err != nil {
		return &QuestionsRepo{}, err
	}

	collection := client.Database(db).Collection("questions")

	return &QuestionsRepo{collection}, nil
}
