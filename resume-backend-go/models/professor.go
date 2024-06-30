package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Professor struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name  string             `bson:"name" json:"name"`
    Sex   string             `bson:"sex" json:"sex"`
    Email string             `bson:"email" json:"email"`
	Students  []primitive.ObjectID `bson:"students,omitempty" json:"students,omitempty"`
}
