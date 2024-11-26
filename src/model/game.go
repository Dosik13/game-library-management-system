package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID              primitive.ObjectID `bson:"_id"`
	Title           string             `bson:"title"`
	Developer       Developer          `bson:"developer"`
	Genre           string             `bson:"genre"`
	PublicationYear int                `bson:"year"`
	Available       bool               `bson:"available"`
}
