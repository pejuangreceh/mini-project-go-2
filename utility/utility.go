package utility

import (
	"crud_api/dto"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
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
func VerifyJWT(receivedToken, secretKey string) (*dto.MyClaimsResponse, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var myClaims dto.MyClaimsResponse

		err := mapstructure.Decode(claims, &myClaims)
		if err != nil {
			return nil, err
		}

		return &myClaims, nil
	}
	return nil, errors.New("Token tidak valid")
}
