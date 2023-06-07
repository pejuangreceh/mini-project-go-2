package utility

import (
	"crud_api/dto"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

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

func GenerateJWT(res dto.ActorDataResponse, key string) (string, error) {
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
