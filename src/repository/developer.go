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

func (r *DeveloperRepository) AddDeveloper(ctx context.Context, developer model.Developer) (*model.Developer, error) {
	developer.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, developer)
	if err != nil {
		return nil, err
	}

	return &developer, nil
}

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
