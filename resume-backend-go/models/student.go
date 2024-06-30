package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name  string             `bson:"name" json:"name"`
    Sex   string             `bson:"sex" json:"sex"`
    Email string             `bson:"email" json:"email"`
	Professors  []primitive.ObjectID `bson:"professors,omitempty" json:"professors,omitempty"`

}
