package utility

import (
	"crud_api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func customersHandlerAuth(c *gin.Context) string {
	// Get the Authorization header value
	tokenString := c.GetHeader("Authorization")

	// Extract the JWT token from the header
	// Assuming the format is "Bearer <token>"
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		// Handle invalid Authorization header
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid Authorization header"})
	}
	token := parts[1]
	return token

}
