package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskModel struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt time.Time
	Completed bool `json:"completed"`
}
