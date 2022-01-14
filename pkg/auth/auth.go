package auth

import (
	"errors"
	"os"
	"time"

	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	issuer        = "go-back"
	audience      = "go-back-client"
	tokenDuration = time.Hour * 24
)

var (
	secret          = os.Getenv("JWT_SECRET")
	ErrInvalidToken = errors.New("invalid token")
	ErrNoPermission = errors.New("no permission")
)

type UserClaims struct {
	jwt.StandardClaims
	UserID string `json:"uid"`
}

func CreateToken(userID string) (string, error) {
	// https://auth0.com/docs/security/tokens/json-web-tokens/json-web-token-claims
	now := time.Now()

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  audience,
			ExpiresAt: now.Add(tokenDuration).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  now.Unix(),
			Issuer:    issuer,
			NotBefore: now.Unix(),
			Subject:   userID,
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authorize(tokenString string, userID string) error {
	claims, err := verifyToken(tokenString)
	if err != nil {
		return ErrInvalidToken
	}
	if claims.UserID != userID {
		return ErrNoPermission
	}

	return nil
}

func verifyToken(tokenString string) (*UserClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Error("unexpected token signing method")
			return nil, errors.New("unexpected token signing method")
		}

		return []byte(secret), nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, keyFunc)
	if err != nil {
		log.Error("failed to parse token", logField.String("token", tokenString), logField.Error(err))
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		log.Error("invalid token claims")
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
