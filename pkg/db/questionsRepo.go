package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/undeadops/enigma/pkg/models"
)

// SaveResponse - Save Reponse to database
func (qr *QuestionsRepo) SaveResponse(ctx context.Context, r *models.Response) error {

	insertResult, err := qr.InsertOne(ctx, r)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted: %v", insertResult.InsertedID)
	return nil
}

// DeleteResponse - Delete a Response
func (qr *QuestionsRepo) DeleteResponse(ctx context.Context, id string) error {
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

// ListResponses - Return a List of Responses
func (qr *QuestionsRepo) ListResponses(ctx context.Context) ([]*models.Response, error) {
	findOptions := options.Find()
	findOptions.SetLimit(50)

	var results []*models.Response

	cur, err := qr.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return results, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem models.Response
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Printf("Error Decoding Element: %v", err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		fmt.Printf("Error in Cursor: %v\n", err)
	}

	// Close the cursor once finished
	cur.Close(ctx)
	return results, nil
}

// GetByID - Return a reponse based ID given
func (qr *QuestionsRepo) GetByID(ctx context.Context, id string) (*models.Response, error) {
	var result models.Response

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.Response{}, err
	}

	filter := bson.M{"_id": objID}
	err = qr.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &models.Response{}, err
	}

	return &result, nil
}
