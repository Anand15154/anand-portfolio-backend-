package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title" binding:"required"`
	Description  string             `json:"description" bson:"description" binding:"required"`
	ImageURL     string             `json:"image_url" bson:"image_url"`
	LiveURL      string             `json:"live_url" bson:"live_url"`
	GitHubURL    string             `json:"github_url" bson:"github_url"`
	Technologies []string           `json:"technologies" bson:"technologies"`
	Category     string             `json:"category" bson:"category"`
	Featured     bool               `json:"featured" bson:"featured"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type ProjectResponse struct {
	ID           primitive.ObjectID `json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	ImageURL     string             `json:"image_url"`
	LiveURL      string             `json:"live_url"`
	GitHubURL    string             `json:"github_url"`
	Technologies []string           `json:"technologies"`
	Category     string             `json:"category"`
	Featured     bool               `json:"featured"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
