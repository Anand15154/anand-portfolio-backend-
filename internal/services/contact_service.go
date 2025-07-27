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

type ContactService struct {
	db           *database.MongoDB
	collection   *mongo.Collection
	emailService *EmailService
}

func NewContactService(db *database.MongoDB, emailService *EmailService) *ContactService {
	return &ContactService{
		db:           db,
		collection:   db.GetCollection("contacts"),
		emailService: emailService,
	}
}

func (s *ContactService) CreateContact(contact *models.Contact) error {
	contact.CreatedAt = time.Now()
	contact.Read = false

	_, err := s.collection.InsertOne(context.Background(), contact)
	if err != nil {
		return err
	}

	// Send email notification
	if s.emailService != nil {
		go func() {
			if err := s.emailService.SendContactNotification(
				contact.Name,
				contact.Email,
				contact.Subject,
				contact.Message,
			); err != nil {
				// Log error but don't fail the request
				// In production, you might want to use a proper logger
				println("Failed to send email notification:", err.Error())
			}
		}()
	}

	return nil
}

func (s *ContactService) GetAllContacts() ([]models.ContactResponse, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var contacts []models.ContactResponse
	if err = cursor.All(context.Background(), &contacts); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *ContactService) GetContactByID(id string) (*models.ContactResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var contact models.ContactResponse
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&contact)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (s *ContactService) MarkAsRead(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{"read": true}},
	)
	return err
}

func (s *ContactService) DeleteContact(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
