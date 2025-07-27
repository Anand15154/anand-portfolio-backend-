package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/models"
)

type ProjectService struct {
	db         *database.MongoDB
	collection *mongo.Collection
}

func NewProjectService(db *database.MongoDB) *ProjectService {
	return &ProjectService{
		db:         db,
		collection: db.GetCollection("projects"),
	}
}

func (s *ProjectService) CreateProject(project *models.Project) error {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	_, err := s.collection.InsertOne(context.Background(), project)
	return err
}

func (s *ProjectService) GetAllProjects() ([]models.ProjectResponse, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var projects []models.ProjectResponse
	if err = cursor.All(context.Background(), &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) GetFeaturedProjects() ([]models.ProjectResponse, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.collection.Find(context.Background(), bson.M{"featured": true}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var projects []models.ProjectResponse
	if err = cursor.All(context.Background(), &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) GetProjectByID(id string) (*models.ProjectResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var project models.ProjectResponse
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectService) UpdateProject(id string, project *models.Project) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	project.UpdatedAt = time.Now()
	project.ID = objectID

	_, err = s.collection.ReplaceOne(
		context.Background(),
		bson.M{"_id": objectID},
		project,
	)
	return err
}

func (s *ProjectService) DeleteProject(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
