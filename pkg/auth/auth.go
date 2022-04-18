package auth

import (
	"log"
	"time"

	cfg "github.com/cave/config"
	"github.com/cave/pkg/models"
	"github.com/cave/pkg/utils"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// RefreshClaims represents refresh token JWT claims
type RefreshClaims struct {
	RefreshTokenID string `json:"refreshTokenID"`
	ExternalID     string `json:"userID"`
	Role           string `json:"role"`
	jwt.StandardClaims
}

// AccessClaims represents access token JWT claims
type AccessClaims struct {
	AccessTokenID string `json:"accessTokenID"`
	ExternalID    string `json:"userID"`
	Role          string `json:"role"`
	jwt.StandardClaims
}

// IssueAccessToken generate access tokens used for authentication
func IssueAccessToken(u models.User) (string, error) {
	expireTime := time.Now().Add(time.Hour) // 1 hour

	// Generate encoded token
	claims := AccessClaims{
		uuid.New().String(),
		u.Id.Hex(),
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(cfg.GetConfig().JWTAccessSecret))
}

// IssueRefreshToken generate refresh tokens used for refreshing authentication
func IssueRefreshToken(u models.User) (string, error) {
	expireTime := time.Now().Add((24 * time.Hour) * 14) // 14 days

	// Generate encoded token
	claims := RefreshClaims{
		uuid.New().String(),
		u.Id.Hex(),
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(cfg.GetConfig().JWTRefreshSecret))
}

// IssueRefreshToken generate verification token used for email verification
func IssueVerificationToken(u models.User) (string, error) {
	expireTime := time.Now().Add((24 * time.Hour) * 7) // 7 days

	// Generate encoded token
	claims := RefreshClaims{
		uuid.New().String(),
		u.Id.Hex(),
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(cfg.GetConfig().JWTRefreshSecret))
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	hmacSecret := []byte(cfg.GetConfig().JWTRefreshSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, utils.ErrInvalidAuthToken
	}
}
