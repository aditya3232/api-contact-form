// Package handlers contains the HTTP handler implementations for various endpoint
//
// Specifically, the HealthHandler provides a health check endpoint to verify that the API is running correctly.
//
// Author: Adit
package handlers

import "github.com/gin-gonic/gin"

// HealthHandler handles HTTP requests related to the health checks.
type HealthHandler struct{}

// NewHealthHandler creates a new instances of HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck responds with a simple message indicating that the API is running.
//
// It returns a JSON reponse with a 200 status code and a message
// conforming the operational status of the API.
func (*HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API is up and running",
		Data:    nil,
	})
}
