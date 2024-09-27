package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Password []byte      `json:"password"`
}