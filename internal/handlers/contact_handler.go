package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
)

type ContactHandler struct {
	contactService *services.ContactService
}

func NewContactHandler(contactService *services.ContactService) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
	}
}

// CreateContact handles contact form submissions
func (h *ContactHandler) CreateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contact message sent successfully",
		"contact": contact,
	})
}

// GetAllContacts retrieves all contact messages (admin only)
func (h *ContactHandler) GetAllContacts(c *gin.Context) {
	contacts, err := h.contactService.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
	})
}

// GetContactByID retrieves a specific contact message
func (h *ContactHandler) GetContactByID(c *gin.Context) {
	id := c.Param("id")
	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})
}

// MarkAsRead marks a contact message as read
func (h *ContactHandler) MarkAsRead(c *gin.Context) {
	id := c.Param("id")
	if err := h.contactService.MarkAsRead(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark contact as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact marked as read",
	})
}

// DeleteContact deletes a contact message
func (h *ContactHandler) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if err := h.contactService.DeleteContact(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact deleted successfully",
	})
}
