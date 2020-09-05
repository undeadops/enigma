package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuestionsRepo - MonogoDB Connection
type QuestionsRepo struct {
	*mongo.Collection
}

// QuestionSetRepo - MongoDB Collection
type QuestionSetRepo struct {
	*mongo.Collection
}

func connectLoop(ctx context.Context, client *options.ClientOptions) (*mongo.Client, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := 5 * time.Minute

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := mongo.Connect(ctx, client)
			if err == nil {
				return db, nil
			}
			log.Println(errors.Wrapf(err, "Ticker Failed to connect to db %s", client.Hosts))
		}
	}
}

// NewQuestionsRepo - Connect to Database and return connection
func NewQuestionsRepo(ctx context.Context, config string, db string) (*QuestionsRepo, error) {
	clientOptions := options.Client().ApplyURI(config)

	// TODO: Implement Initial Retry Logic Here Maybe? or higherlevel in main function?
	// Connect to MongoDB
	//client, err := mongo.Connect(ctx, clientOptions)
	client, err := connectLoop(ctx, clientOptions)

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

// NewQuestionSetRepo - Connect to Database and return connection
func NewQuestionSetRepo(ctx context.Context, config string, db string) (*QuestionSetRepo, error) {
	clientOptions := options.Client().ApplyURI(config)

	// TODO: Implement Initial Retry Logic Here Maybe? or higherlevel in main function?
	// Connect to MongoDB
	//client, err := mongo.Connect(ctx, clientOptions)
	client, err := connectLoop(ctx, clientOptions)

	if err != nil {
		return &QuestionSetRepo{}, err
	}
	// Check the connection - Reduces Client resiliance...
	err = client.Ping(ctx, nil)

	if err != nil {
		return &QuestionSetRepo{}, err
	}

	collection := client.Database(db).Collection("question_set")

	return &QuestionSetRepo{collection}, nil
}
