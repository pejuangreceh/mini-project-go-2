package account

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
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
type RegisterStatus struct {
	AdminID uint8  `json:"admin_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
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
func (RegisterStatus) TableName() string {
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

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GenerateJWT(res ActorDataResponse, key string) (string, error) {
	claims := jwt.MapClaims{
		"id":          res.ID,
		"username":    res.Username,
		"role_id":     res.RoleID,
		"is_verified": res.IsVerified,
		"is_active":   res.IsActive,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour * 1).Unix(),
	}
	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "Terjadi Kesalahan", err
	}
	return signedToken, nil
	// Gunakan signedToken seperti yang Anda butuhkan
	//fmt.Println(signedToken)
}
