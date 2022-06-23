package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogPost struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string
	Author         string
	CreateDateTime string `bson:"createDateTime,omitempty"`
	UpdateDateTime string `bson:"updateDateTime,omitempty"`
	Body           []string
	Tags           []string
}
