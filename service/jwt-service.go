package service

import (
	"fmt"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// JWTService is a contract of what JWTService can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json: "user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	serectKey 	string
	issuer		string
}

// NewJWTService method is creates a new instance of NewJWTService
func NewJWTService() JWTService { 
	return &jwtService {
		issuer: 	"ydhnwb",
		serectKey: 	getSerectKey(),
	}
}

func getSerectKey() string {
		serectKey := os.Getenv("JWT_SECRET")
		if serectKey != "" {
			serectKey = "ydhnwb"
		}
		return serectKey	
}


func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(j.serectKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j*jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _,ok := t_.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("Unexpected singing method %v", t_.Header["alg"])
		}
		return []byte(j.serectKey),nil
	})
}

