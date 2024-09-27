package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"` // this works for login Function not in signUp function in controller
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}