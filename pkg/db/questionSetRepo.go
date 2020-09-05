package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/undeadops/enigma/pkg/models"
)

// SaveQuestionSet - Save Set of Questions to Database
func (qsr *QuestionSetRepo) SaveQuestionSet(ctx context.Context, r *models.QuestionSet) error {
	// Set ID
	r.ID = primitive.NewObjectID()
	// Set time
	r.Date = time.Now()

	insertResult, err := qsr.InsertOne(ctx, r)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted: %v", insertResult.InsertedID)
	return nil
}

// DeleteQuestionSet - Delete a Set of Questions
func (qr *QuestionSetRepo) DeleteQuestionSet(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	_, err = qr.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

// ListQuestionSet - Return a List of Questions
func (qsr *QuestionSetRepo) ListQuestionSet(ctx context.Context) ([]*models.QuestionSet, error) {
	findOptions := options.Find()
	findOptions.SetLimit(50)

	var results []*models.QuestionSet

	cur, err := qsr.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return results, err
	}

	// Close the cursor once finished
	defer cur.Close(ctx)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem models.QuestionSet
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Printf("Error Decoding Element: %v", err)
		}

		results = append(results, &elem)
	}

	// if err := cur.Err(); err != nil {
	// 	fmt.Printf("Error in Cursor: %v\n", err)
	// }

	return results, nil
}

// GetQuestionSet - Return a reponse based ID given
func (qsr *QuestionSetRepo) GetQuestionSet(ctx context.Context, id string) (*models.QuestionSet, error) {
	var result models.QuestionSet

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.QuestionSet{}, err
	}

	filter := bson.M{"_id": objID}
	err = qsr.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &models.QuestionSet{}, err
	}

	return &result, nil
}
