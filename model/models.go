package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movies struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"movie,omitempty"`
	Year    int                `json:"year,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
