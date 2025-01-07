package controllers

import (
	"jwt-golang/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// EmailPayload represents the payload expected for sending an email
type EmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

// SendEmail handles sending an email
func SendEmail(c *fiber.Ctx) error {
	// Parse request body into EmailPayload struct
	var payload EmailPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate payload fields
	if payload.Recipient == "" || payload.Subject == "" || payload.Message == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "All fields (recipient, subject, message) are required",
		})
	}

	// Use the helpers.SendEmail function to send the email
	err := helpers.SendEmail(payload.Recipient, payload.Subject, payload.Message)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to send email",
			"details": err.Error(),
		})
	}

	// Respond with success
	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
	})
}
