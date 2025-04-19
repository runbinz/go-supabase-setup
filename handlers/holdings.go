package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/runbinz/dashboard/middleware"
	"github.com/runbinz/dashboard/services"
)

// Package handlers provides HTTP handlers for the application.
// These handlers process incoming requests and return responses.
// They act as the bridge between the HTTP layer and the business logic, handling input validation and response formatting.

// GetHoldings handles GET /api/get-holdings
// It retrieves the user ID from the context and fetches the user's holdings.
// If the user is not authenticated, it returns an unauthorized error.
// If fetching holdings fails, it returns an internal server error.
// This handler is responsible for returning the user's investment holdings in response to a client request.
func GetHoldings(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	holdings, err := services.GetUserHoldings(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch holdings"})
		return
	}

	c.JSON(http.StatusOK, holdings)
}
