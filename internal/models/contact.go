package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Subject   string             `json:"subject" bson:"subject" binding:"required"`
	Message   string             `json:"message" bson:"message" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Read      bool               `json:"read" bson:"read"`
}

type ContactResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Subject   string             `json:"subject"`
	Message   string             `json:"message"`
	CreatedAt time.Time          `json:"created_at"`
	Read      bool               `json:"read"`
}
