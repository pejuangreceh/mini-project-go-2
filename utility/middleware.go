package utility

import (
	"crud_api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AdminAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		// Handle invalid Authorization header
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid Authorization header"})
		c.Abort() // Abort further execution
		return
	}
	token := parts[1]
	tokenData, err := VerifyJWT(token, "rahasia")
	if err != nil {
		//c.JSON(http.StatusConflict, dto.ErrorResponse{Error: ("Invalid Data " + token + err.Error())})
		c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "Invalid Token"})
		c.Abort() // Abort further execution
		return
	}
	if !(tokenData.IsActive == "true" && tokenData.RoleID < 3) {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "admin masih tidak aktif"})
		c.Abort() // Abort further execution
		return
	}
}

func SuperAdminAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		// Handle invalid Authorization header
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid Authorization header"})
		c.Abort() // Abort further execution
		return
	}
	token := parts[1]
	tokenData, err := VerifyJWT(token, "rahasia")
	if err != nil {
		//c.JSON(http.StatusConflict, dto.ErrorResponse{Error: ("Invalid Data " + token + err.Error())})
		c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "Invalid Token"})
		c.Abort() // Abort further execution
		return
	}
	if !(tokenData.IsActive == "true" && tokenData.RoleID == 1) {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "anda bukan superadmin"})
		c.Abort() // Abort further execution
		return
	}
}
