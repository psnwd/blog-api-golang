package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId string, isAdmin bool, email string, password string) (t string, err error)
	ParseToken(tokenString string) (claims JwtCustomClaim, err error)
}

type JwtCustomClaim struct {
	UserID   string
	IsAdmin  bool
	Email    string
	Password string
	jwt.StandardClaims
}

type jwtService struct {
	secretKey []byte
	issuer    string
	expired   time.Duration
}

func NewJWTService(secretKey, issuer string, expired time.Duration) JWTService {
	return &jwtService{
		issuer:    issuer,
		secretKey: []byte(secretKey),
		expired:   expired,
	}
}

func (j *jwtService) GenerateToken(userID string, isAdmin bool, email string, password string) (t string, err error) {
	claims := &JwtCustomClaim{
		UserID:   userID,
		IsAdmin:  isAdmin,
		Email:    email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expired).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString(j.secretKey)
	return
}

func (j *jwtService) ParseToken(tokenString string) (claims JwtCustomClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil || !token.Valid {
		return JwtCustomClaim{}, errors.New("token is not valid")
	}

	if customClaims, ok := token.Claims.(*JwtCustomClaim); ok {
		return *customClaims, nil
	}

	return JwtCustomClaim{}, errors.New("failed to parse token claims")
}
