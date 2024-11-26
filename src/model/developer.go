package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Developer struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	MainHq string             `bson:"mainhq"`
}
