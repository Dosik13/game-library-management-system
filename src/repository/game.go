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

type GameRepository struct {
	collection *mongo.Collection
}

func NewGameRepository(URI, dbName string) (_interface.GameRepositorer, error) {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &GameRepository{
		collection: db.Collection("games"),
	}, nil
}

func (r *GameRepository) GetAllGames(ctx context.Context) ([]model.Game, error) {
	var gs []model.Game

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &gs); err != nil {
		return nil, err
	}

	return gs, nil
}

func (r *GameRepository) GetGameById(ctx context.Context, id string) (*model.Game, error) {
	var game model.Game
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": i}).Decode(&game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (r *GameRepository) AddGame(ctx context.Context, game model.Game) (*model.Game, error) {
	var developer model.Developer
	err := r.collection.Database().Collection("developers").FindOne(ctx, bson.M{"_id": game.Developer.ID}).Decode(&developer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("developer does not exist")
		}
		return nil, err
	}
	game.ID = primitive.NewObjectID()
	_, err = r.collection.InsertOne(ctx, game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (r *GameRepository) UpdateAvailability(ctx context.Context, id string) (*model.Game, error) {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var game model.Game
	err = r.collection.FindOne(ctx, bson.M{"_id": i}).Decode(&game)
	if err != nil {
		return nil, err
	}
	available := !game.Available
	_, err = r.collection.UpdateByID(ctx, i, bson.M{"$set": bson.M{"available": available}})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *GameRepository) DeleteGame(ctx context.Context, id string) error {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": i})
	if err != nil {
		return err
	}
	return nil
}

func (r *GameRepository) FindGamesByDeveloper(ctx context.Context, developerName string) ([]model.Game, error) {
	var developer model.Developer
	err := r.collection.Database().Collection("developers").FindOne(ctx, bson.M{"name": developerName}).Decode(&developer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("developer does not exist")
		}
		return nil, err
	}

	var games []model.Game

	cursor, err := r.collection.Find(ctx, bson.M{"developer._id": developer.ID})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *GameRepository) DeleteManyGamesByDeveloper(ctx context.Context, developerId string) error {
	id, err := primitive.ObjectIDFromHex(developerId)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteMany(ctx, bson.M{"developer._id": id})
	if err != nil {
		return err
	}
	return nil
}