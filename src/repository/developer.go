package repository

import (
	"context"
	"errors"
	"game-library-management-system/src/interface"
	"game-library-management-system/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeveloperRepository struct {
	collection *mongo.Collection
}

// NewDeveloperRepository creates a new DeveloperRepository instance.
// Connects to the MongoDB database using the provided URI and database name.
// Returns the DeveloperRepositorer interface or an error if the connection fails.
func NewDeveloperRepository(URI, dbName string) (_interface.DeveloperRepositorer, error) {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &DeveloperRepository{
		collection: db.Collection("developers"),
	}, nil
}

// GetAllDevelopers retrieves all developers from the collection.
// Takes a context for managing request lifetime.
// Returns a slice of Developer models or an error if the operation fails.
func (r *DeveloperRepository) GetAllDevelopers(ctx context.Context) ([]model.Developer, error) {
	var devs []model.Developer

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &devs); err != nil {
		return nil, err
	}

	return devs, nil
}

// GetDeveloperById retrieves a developer by their ID from the collection.
// Takes a context for managing request lifetime and the developer ID as a string.
// Returns a Developer model or an error if the operation fails.
func (r *DeveloperRepository) GetDeveloperById(ctx context.Context, id string) (*model.Developer, error) {
	var dev model.Developer
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": i}).Decode(&dev)
	if err != nil {
		return nil, err
	}

	return &dev, nil
}

// AddDeveloper inserts a new developer into the collection.
// Takes a context for managing request lifetime and a Developer model.
// Returns the inserted Developer model or an error if the operation fails.
func (r *DeveloperRepository) AddDeveloper(ctx context.Context, developer model.Developer) (*model.Developer, error) {
	developer.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, developer)
	if err != nil {
		return nil, err
	}

	return &developer, nil
}

// UpdateDeveloper updates an existing developer in the collection.
// Takes a context for managing request lifetime, the developer ID as a string, and a Developer model.
// Returns the updated Developer model or an error if the operation fails.
func (r *DeveloperRepository) UpdateDeveloper(ctx context.Context, id string, developer model.Developer) (*model.Developer, error) {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	developer.ID = i
	result, err := r.collection.UpdateByID(ctx, i, bson.M{"$set": developer})
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("developer id not found")
	}
	return nil, nil
}

// DeleteDeveloper removes a developer from the collection by their ID.
// Takes a context for managing request lifetime and the developer ID as a string.
// Returns an error if the operation fails or if no document is found.
func (r *DeveloperRepository) DeleteDeveloper(ctx context.Context, id string) error {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	info, err := r.collection.DeleteOne(ctx, bson.M{"_id": i})
	if err != nil {
		return err
	}
	if info.DeletedCount == 0 {
		return errors.New("no document found")
	}
	return nil
}
