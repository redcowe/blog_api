package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogPost struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string
	Author          string
	CreatedDateTime string
	UpdateDateTime  string
	Body            string
	Tags            []string
}
