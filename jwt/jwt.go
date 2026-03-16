package jwt

import (
	"auth/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Type     string `json:"type"` // access or refresh
	jwt.RegisteredClaims
}

func GenerateAccessToken(cfg config.Config, userID int64, username, email string) (string, time.Time, error) {

	expiry := time.Now().Add(time.Duration(cfg.JWTAccessExpiryTimeMinute) * time.Minute)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		Type:     "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.JWTAccessSecrect))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("token signing failed : %w", err)
	}

	return signedToken, expiry, nil
}

func ValidateAccessToken(token string, cfg *config.Config) (*Claims, error) {
	claims, err := validatetoken(token, cfg.JWTAccessSecrect, "access")
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// pirvate function
func validatetoken(tokenString, secret, tokenType string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Type != tokenType {
		return nil, fmt.Errorf("invalid token type")
	}

	return claims, nil
}
