package account

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Actors struct {
	gorm.Model
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RoleID     uint8  `json:"role_id" binding:"required"`
	IsVerified string `json:"is_verified" binding:"required"`
	IsActive   string `json:"is_active" binding:"required"`
}

type Approval struct {
	gorm.Model
	IsVerified string `json:"is_verified" binding:"required"`
}

type Activate struct {
	gorm.Model
	IsActive string `json:"is_active" binding:"required"`
}

type Register struct {
	AdminID      uint8  `json:"admin_id" binding:"required"`
	SuperAdminID uint8  `json:"superadmin_id" binding:"required"`
	Status       string `json:"status"`
}

func (Approval) TableName() string {
	return "actors" // specify the actual table name here
}
func (Activate) TableName() string {
	return "actors" // specify the actual table name here
}

func (Register) TableName() string {
	return "register" // specify the actual table name here
}
func HashPassword(password string) (string, error) {
	// Generate the hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Convert the hashed password byte slice to a string
	hashedPassword := string(hash)
	return hashedPassword, nil
}
func CheckPassword(password, hashedPassword string) error {
	// Compare the password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
