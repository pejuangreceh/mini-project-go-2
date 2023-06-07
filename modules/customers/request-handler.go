package customers

import (
	"crud_api/dto"
	"crud_api/entities"
	"crud_api/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type RequestHandler struct {
	ctrl *Controller
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type CreateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Read(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		// Handle invalid Authorization header
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid Authorization header"})
		return
	}
	token := parts[1]
	tokenData, err := utility.VerifyJWT(token, "rahasia")
	if err != nil {
		//c.JSON(http.StatusConflict, dto.ErrorResponse{Error: ("Invalid Data " + token + err.Error())})
		c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "Invalid Token"})
		return
	}
	if !(tokenData.IsActive == "true" && tokenData.RoleID < 3) {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "admin masih tidak aktif"})
	} else {
		res, err := h.ctrl.Read()
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}

}
func (h RequestHandler) ReadID(c *gin.Context) {
	customerID := c.Param("id")
	res, err := h.ctrl.ReadID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) Update(c *gin.Context) {
	customer := entities.Customers{}
	customerID := c.Param("id")
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(customer)
	res, err := h.ctrl.Update(customer, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) Delete(c *gin.Context) {
	customerID := c.Param("id")
	res, err := h.ctrl.Delete(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
